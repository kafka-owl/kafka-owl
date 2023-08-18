/**
 * Copyright 2022 Redpanda Data, Inc.
 *
 * Use of this software is governed by the Business Source License
 * included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
 *
 * As of the Change Date specified in that file, in accordance with
 * the Business Source License, use of this software will be governed
 * by the Apache License, Version 2.0
 */

import React, { Component } from 'react';
import { observer } from 'mobx-react';
import { ConfigProvider, Table } from 'antd';
import { makePaginationConfig, renderLogDirSummary, sortField, WarningToolip } from '../../misc/common';
import { Partition, PartitionReassignmentsPartition, Topic } from '../../../state/restInterfaces';
import { BrokerList } from '../../misc/BrokerList';
import { IndeterminateCheckbox } from './components/IndeterminateCheckbox';
import { SelectionInfoBar } from './components/StatisticsBar';
import { prettyBytesOrNA } from '../../../utils/utils';
import { ColumnProps } from 'antd/lib/table/Column';
import { DefaultSkeleton, ZeroSizeWrapper, InfoText } from '../../../utils/tsxUtils';
import { api } from '../../../state/backendApi';
import { computed, IReactionDisposer, makeObservable, observable, transaction } from 'mobx';
import { PartitionSelection } from './ReassignPartitions';
import Highlighter from 'react-highlight-words';
import { uiSettings } from '../../../state/ui';
import { ColumnFilterItem, ColumnsType, ExpandableConfig, TableRowSelection } from 'antd/lib/table/interface';
import { SearchOutlined, WarningTwoTone } from '@ant-design/icons';
import { SearchTitle } from '../../misc/KowlTable';
import { Popover } from '@redpanda-data/ui'

export type TopicWithPartitions = Topic & { partitions: Partition[], activeReassignments: PartitionReassignmentsPartition[] };

@observer
export class StepSelectPartitions extends Component<{ partitionSelection: PartitionSelection, throttledTopics: string[] }> {
    pageConfig = makePaginationConfig(uiSettings.reassignment.pageSizeSelect ?? 10);
    autorunHandle: IReactionDisposer | undefined = undefined;
    @observable filterOpen = false; // topic name searchbar

    @observable selectedBrokerFilters: (string | number)[] | null = null;
    expandableConfig: ExpandableConfig<TopicWithPartitions>;

    constructor(props: any) {
        super(props);
        this.setSelection = this.setSelection.bind(this);
        this.setTopicSelection = this.setTopicSelection.bind(this);
        this.isSelected = this.isSelected.bind(this);
        this.getSelectedPartitions = this.getSelectedPartitions.bind(this);
        this.getTopicCheckState = this.getTopicCheckState.bind(this);
        this.getRowKey = this.getRowKey.bind(this);

        this.expandableConfig = {
            // expandedRowKeys: this.expandedTopics.slice(),
            // onExpandedRowsChange: keys => {
            //     // console.log('replacing expanded row keys', { current: this.expandedTopics, new: keys })
            //     this.expandedTopics = keys as string[];
            // },
            // expandIconColumnIndex: 1,
            expandedRowRender: topic => topic.partitions
                ? <SelectPartitionTable
                    topic={topic}
                    topicPartitions={topic.partitions}
                    isSelected={this.isSelected}
                    setSelection={this.setSelection}
                    getSelectedPartitions={() => this.getSelectedPartitions(topic.topicName)}
                />
                : <>Error loading partitions</>,
        };

        makeObservable(this);
    }

    componentWillUnmount() {
        if (this.autorunHandle) {
            this.autorunHandle();
            this.autorunHandle = undefined;
        }
    }

    render() {
        if (!api.topics) return DefaultSkeleton;

        const query = uiSettings.reassignment.quickSearch ?? '';
        const filterActive = query.length > 1;

        let searchRegex: RegExp | undefined = undefined;
        if (filterActive) try { searchRegex = new RegExp(uiSettings.reassignment.quickSearch, 'i') } catch { return null; }

        const rackToBrokers = this.rackToBrokers;

        const columns: ColumnProps<TopicWithPartitions>[] = [
            {
                title: <SearchTitle title="Topic" observableFilterOpen={this} observableSettings={uiSettings.reassignment} />,
                dataIndex: 'topicName',
                render: (_v, record) => {
                    const content = filterActive
                        ? <Highlighter searchWords={[query]} textToHighlight={record.topicName} />
                        : record.topicName;

                    if (this.props.throttledTopics.includes(record.topicName))
                        return <div>
                            <span>{content}</span>
                            <WarningToolip content="Topic replication is throttled" position="top" />
                        </div>

                    return content;
                },

                sorter: sortField('topicName'), defaultSortOrder: 'ascend',

                // to support both filters at the same time (topic and brokers), both filters *must* be in controlled mode
                filteredValue: filterActive ? [query] : undefined,
                onFilter: (_value, record) => searchRegex?.test(record.topicName) ?? false,

                filterIcon: filterIcon(filterActive),
                filterDropdown: <></>,
                filterDropdownVisible: this.filterOpen,
                onFilterDropdownVisibleChange: visible => {
                    // only accept requests to open the filter
                    if (visible) this.filterOpen = visible;
                },
            },
            {
                title: 'Partitions', width: 140,
                render: (_, t) => {
                    const errors = t.partitions.count(p => p.hasErrors);
                    if (errors == 0) return t.partitionCount;

                    return <span>
                        <span>{t.partitionCount - errors} / {t.partitionCount}</span>
                        {' '}
                        {renderPartitionErrorsForTopic(errors)}
                    </span>
                },
                sorter: sortField('partitionCount')
            },
            {
                title: 'Replication Factor', width: 160, render: (_t, r) => {
                    if (r.activeReassignments.length == 0) return r.replicationFactor;
                    return <InfoText tooltip="While reassignment is active, replication factor is temporarily doubled." maxWidth="180px">
                        {r.replicationFactor}
                    </InfoText>
                },
                sorter: sortField('replicationFactor')
            },
            {
                title: 'Brokers', width: 160, dataIndex: 'partitions',
                render: (_value, record) => record.partitions?.map(p => p.leader).distinct().length ?? 'N/A',

                // to support both filters at the same time (topic and brokers), both filters *must* be in controlled mode
                filteredValue: this.selectedBrokerFilters,
                filters: this.brokerFilters,
                onFilter: (v, r) => {
                    if (typeof v === 'number') {
                        // Broker ID
                        return r.partitions.any(p => p.replicas.includes(v));
                    }
                    if (typeof v === 'string') {
                        // Rack
                        const brokers = rackToBrokers.get(v);
                        if (!brokers) return false;
                        return r.partitions.any(p => p.replicas.intersection(brokers).length > 0);
                    }
                    return false;
                },
                filterMultiple: true,
            },
            {
                title: 'Size', width: 110,
                render: (_v, r) => renderLogDirSummary(r.logDirSummary),
                sorter: (a, b) => a.logDirSummary.totalSizeBytes - b.logDirSummary.totalSizeBytes
            },
        ];


        return <>
            <div style={{ margin: '1em 1em 2em 1em' }}>

                {/* Current Selection */}
                <SelectionInfoBar partitionSelection={this.props.partitionSelection} margin="2em 0em 1em 0.3em" />

                {/* Topic / Partitions */}
                <ConfigProvider getPopupContainer={t => {
                    let c = t as HTMLElement | null | undefined;
                    while (c && c.tagName != 'TH' && c.tagName != 'TD')
                        c = c.parentElement;
                    return c ?? t ?? document.body;
                }}>

                    <Table
                        style={{ margin: '0', }} size={'middle'}
                        pagination={this.pageConfig}
                        onChange={(p, filters, _sorters) => {
                            if (p.pageSize) uiSettings.reassignment.pageSizeSelect = p.pageSize;
                            this.pageConfig.current = p.current;
                            this.pageConfig.pageSize = p.pageSize;

                            const brokerFilters = filters['partitions']?.filterNull() ?? null;
                            brokerFilters?.removeAll(x => typeof x === 'boolean');

                            this.selectedBrokerFilters = brokerFilters as ((string | number)[] | null);
                        }}

                        dataSource={this.topicPartitions}
                        columns={columns}
                        showSorterTooltip={false}

                        rowKey={this.getRowKey}
                        rowClassName={'pureDisplayRow'}
                        onRow={r => ({
                            onClick: () => this.setTopicSelection(r, !this.getTopicCheckState(r.topicName).checked),
                            // onDoubleClick: () => this.expandedTopics.includes(record.topicName)
                            //     ? this.expandedTopics.remove(record.topicName)
                            //     : this.expandedTopics.push(record.topicName),
                        })}

                        rowSelection={{
                            type: 'checkbox',
                            columnTitle: <div style={{ display: 'flex' }} >
                                <InfoText tooltip={<>
                                    If you want to select multiple adjacent items, you can use the SHIFT key.<br />
                                    Shift-Click selects the first item, last item and all items in between.
                                </>} iconSize="16px" placement="right" />
                            </div>,
                            renderCell: (_value: boolean, record, _index, originNode: React.ReactNode) => {
                                return <IndeterminateCheckbox
                                    originalCheckbox={originNode}
                                    getCheckState={() => this.getTopicCheckState(record.topicName)}
                                />
                            },
                            onSelect: (record, selected: boolean, _selectedRows) => {
                                if (!record.partitions) return;
                                this.setTopicSelection(record, selected);
                            },
                            onSelectMultiple: (selected, _selectedRows, changeRows) => {
                                transaction(() => {
                                    for (const r of changeRows)
                                        for (const p of r.partitions)
                                            this.setSelection(r.topicName, p.id, selected && !p.hasErrors);
                                });
                            },
                        }}

                        expandable={this.expandableConfig}
                    />
                </ConfigProvider>
            </div>
        </>
    }

    getRowKey(r: TopicWithPartitions) {
        return r.topicName;
    }


    setTopicSelection(topic: TopicWithPartitions, isSelected: boolean) {
        transaction(() => {
            for (const p of topic.partitions) {
                const selected = isSelected && !p.hasErrors;
                this.setSelection(topic.topicName, p.id, selected);
            }
        });
    }

    setSelection(topic: string, partition: number, isSelected: boolean) {
        const partitions = this.props.partitionSelection[topic] ?? [];

        if (isSelected) {
            partitions.pushDistinct(partition);
        } else {
            partitions.remove(partition);
        }

        if (partitions.length == 0)
            delete this.props.partitionSelection[topic];
        else
            this.props.partitionSelection[topic] = partitions;
    }

    getSelectedPartitions(topic: string) {
        const partitions = this.props.partitionSelection[topic];
        if (!partitions) return [];
        return partitions;
    }

    isSelected(topic: string, partition: number) {
        const partitions = this.props.partitionSelection[topic];
        if (!partitions) return false;
        return partitions.includes(partition);
    }

    getTopicCheckState(topicName: string): { checked: boolean, indeterminate: boolean } {
        const tp = this.topicPartitions.first(t => t.topicName == topicName);
        if (!tp) return { checked: false, indeterminate: false };

        const selected = this.props.partitionSelection[topicName];
        if (!selected) return { checked: false, indeterminate: false };

        if (selected.length == 0)
            return { checked: false, indeterminate: false };

        const validPartitions = tp.partitions.count(x => !x.hasErrors);
        if (validPartitions > 0 && selected.length == validPartitions)
            return { checked: true, indeterminate: false };

        return { checked: false, indeterminate: true };
    }

    @computed get topicPartitions(): TopicWithPartitions[] {
        if (api.topics == null) return [];
        return api.topics.map(topic => {
            return {
                ...topic,
                partitions: api.topicPartitions.get(topic.topicName)!,
                activeReassignments: this.inProgress.get(topic.topicName) ?? [],
            }
        }).filter(t => t.activeReassignments.length == 0);
    }

    @computed get inProgress() {
        const current = api.partitionReassignments ?? [];
        return current.toMap(x => x.topicName, x => x.partitions);
    }

    @computed get allBrokers() {
        const brokerIdsOfTopicPartitions = this.topicPartitions.flatMap(t => t.partitions).flatMap(p => p.replicas).distinct();
        const allBrokers = api.clusterInfo?.brokers;
        if (!allBrokers || brokerIdsOfTopicPartitions.length == 0) return [];

        return allBrokers.filter(b => brokerIdsOfTopicPartitions.includes(b.brokerId));
    }

    @computed get rackToBrokers(): Map<string, number[]> {
        const brokers = this.allBrokers;
        const racks = brokers.map(b => b.rack).filterFalsy().distinct();

        // rack name => brokerIds
        return racks.toMap(
            r => r,
            r => brokers.filter(b => b.rack === r).map(b => b.brokerId).distinct())
    }

    @computed get brokerFilters(): ColumnFilterItem[] | undefined {
        const brokers = this.allBrokers;
        const racks = brokers.map(b => b.rack).filterFalsy().distinct();

        const brokerFilters: ColumnFilterItem[] = [];

        // Individual Brokers
        const brokerItems = brokers.map(b => ({ text: b.address, value: b.brokerId }));
        if (brokerItems.length > 0)
            brokerFilters.push({
                text: 'Brokers', value: 'Brokers',
                children: brokerItems,
            });

        // Racks
        if (racks.length > 0)
            brokerFilters.push({
                text: 'Racks', value: 'Racks',
                children: racks.map(r => ({ text: r, value: r })),
            });

        return brokerFilters;
    }
}

@observer
export class SelectPartitionTable extends Component<{
    topic: Topic;
    topicPartitions: Partition[];
    setSelection: (topic: string, partition: number, isSelected: boolean) => void;
    isSelected: (topic: string, partition: number) => boolean;
    getSelectedPartitions: () => number[];
}> {
    partitionsPageConfig = makePaginationConfig(100, true);
    scroll = { y: '300px' };
    rowClassName = (p: Partition) => p.hasErrors ? 'errorPartition' : '';

    columns: ColumnsType<Partition> = [
        { width: 100, title: 'Partition', dataIndex: 'id', sortOrder: 'ascend', sorter: (a, b) => a.id - b.id },
        {
            width: undefined, title: 'Brokers', render: (_, p) =>
                Boolean(p.replicas)
                    ? <BrokerList brokerIds={p.replicas} leaderId={p.leader} />
                    : renderPartitionError(p)
        },
        {
            width: 100, title: 'Size', render: (_v, p) => prettyBytesOrNA(p.replicaSize),
            sortOrder: 'ascend', sorter: (a, b) => a.replicaSize - b.replicaSize
        },
    ];

    render() {
        return <div style={{ paddingTop: '4px', paddingBottom: '8px', width: 0, minWidth: '100%' }}>
            <Table size="small" className="nestedTable"
                dataSource={this.props.topicPartitions}
                pagination={this.partitionsPageConfig}
                scroll={this.scroll}
                rowClassName={this.rowClassName}
                rowKey={this.getRowKey}
                rowSelection={this.rowSelection}
                columns={this.columns} />
        </div>;
    }

    getRowKey(p: Partition) {
        return p.id;
    }

    @computed get rowSelection(): TableRowSelection<Partition> {
        return {
            type: 'checkbox',
            columnWidth: '43px',
            columnTitle: <></>,
            selectedRowKeys: this.props.getSelectedPartitions().slice(),
            getCheckboxProps: this.getCheckboxProps,
            onSelect: (p, selected) => this.props.setSelection(this.props.topic.topicName, p.id, selected && !p.hasErrors),
        }
    }

    getCheckboxProps(p: Partition) {
        return { disabled: p.hasErrors };
    }
}

function filterIcon(filterActive: boolean) {
    return <div style={{
        background: filterActive ? 'hsl(208deg, 100%, 93%)' : undefined,
        position: 'absolute',
        left: 0, top: 0, right: 0, bottom: 0
    }}>
        <SearchOutlined style={{ color: filterActive ? '#1890ff' : 'hsl(0deg, 0%, 67%)', fontSize: '14px' }} />
    </div>
}

function renderPartitionError(partition: Partition) {
    const txt = [partition.partitionError, partition.waterMarksError].join('\n\n');

    return <Popover
        title="Partition Error"
        placement="right-start" 
        size="auto"
        hideCloseButton
        content={
            <div style={{ maxWidth: '500px', whiteSpace: 'pre-wrap' }}>
                {txt}
            </div>
        }
    >
        <span>
            <ZeroSizeWrapper justifyContent="center" alignItems="center" width="20px" height="18px">
                <span style={{ fontSize: '19px' }}>
                    <WarningTwoTone twoToneColor="orange" />
                </span>
            </ZeroSizeWrapper>
        </span>
    </Popover>
}

function renderPartitionErrorsForTopic(_partitionsWithErrors: number) {
    return <Popover
        title="Partition Error"
        placement="right-start" 
        size="auto"
        hideCloseButton
        content={
            <div style={{ maxWidth: '500px', whiteSpace: 'pre-wrap' }}>
                Some partitions could not be retreived.<br />
                Expand the topic to see which partitions are affected.
            </div>
        }
    >
        <span>
            <ZeroSizeWrapper justifyContent="center" alignItems="center" width="20px" height="18px">
                <span style={{ fontSize: '19px' }}>
                    <WarningTwoTone twoToneColor="orange" />
                </span>
            </ZeroSizeWrapper>
        </span>
    </Popover>
}
