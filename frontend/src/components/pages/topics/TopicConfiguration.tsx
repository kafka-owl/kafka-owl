import { PencilIcon } from '@heroicons/react/solid';
import {
  Alert,
  AlertIcon,
  Box,
  Button,
  Flex,
  FormField,
  Icon,
  Input,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  PasswordInput,
  Popover,
  RadioGroup,
  SearchField,
  Text,
  Tooltip,
  useToast,
} from '@redpanda-data/ui';
import { Observer, observer, useLocalObservable } from 'mobx-react';
import type { FC } from 'react';
import { useState } from 'react';
import { Controller, type SubmitHandler, useForm } from 'react-hook-form';
import type { ConfigEntryExtended } from '../../../state/restInterfaces';
import { formatConfigValue } from '../../../utils/formatters/ConfigValueFormatter';
import { DataSizeSelect, DurationSelect, NumInput, RatioInput } from './CreateTopicModal/CreateTopicModal';
import './TopicConfiguration.scss';
import { MdInfoOutline } from 'react-icons/md';
import { isServerless } from '../../../config';
import { api } from '../../../state/backendApi';
import { SingleSelect } from '../../misc/Select';

type ConfigurationEditorProps = {
  targetTopic: string | null; // topic name, or null if default configs
  entries: ConfigEntryExtended[];
  onForceRefresh: () => void;
};

type Inputs = {
  valueType: 'default' | 'infinite' | 'custom';
  customValue: string | number | undefined | null;
};

const ConfigEditorForm: FC<{
  editedEntry: ConfigEntryExtended;
  onClose: () => void;
  onSuccess: () => void;
  targetTopic: string;
}> = ({ editedEntry, onClose, targetTopic, onSuccess, onError }) => {
  const toast = useToast();
  const [globalError, setGlobalError] = useState<string | null>(null);

  const {
    control,
    handleSubmit,
    formState: { isSubmitting },
    watch,
  } = useForm<Inputs>({
    defaultValues: {
      valueType: editedEntry.isExplicitlySet ? (editedEntry.value === '-1' ? 'infinite' : 'custom') : 'default',
      customValue: editedEntry.isExplicitlySet && editedEntry.value !== '-1' ? editedEntry.value : undefined,
    },
  });

  const hasInfiniteValue = ['BYTE_SIZE', 'DURATION'].includes(editedEntry.frontendFormat);
  const valueTypeOptions = [];
  valueTypeOptions.push({
    label: 'Default',
    value: 'default',
  });
  if (hasInfiniteValue) {
    valueTypeOptions.push({
      label: 'Infinite',
      value: 'infinite',
    });
  }
  valueTypeOptions.push({
    label: 'Custom',
    value: 'custom',
  });

  const onSubmit: SubmitHandler<Inputs> = async ({ valueType, customValue }) => {
    const operation = valueType === 'infinite' || valueType === 'custom' ? 'SET' : 'DELETE';

    let value: number | string = undefined;
    if (valueType === 'infinite') {
      value = -1;
    } else if (valueType === 'custom') {
      value = customValue;
    }

    try {
      await api.changeTopicConfig(targetTopic, [
        {
          key: editedEntry.name,
          op: operation,
          value: operation === 'SET' ? String(value) : undefined,
        },
      ]);
      toast({
        status: 'success',
        description: (
          <span>
            Successfully updated config <code>{editedEntry.name}</code>
          </span>
        ),
      });
      onSuccess();
      onClose();
    } catch (err) {
      console.error('error while applying config change', { err, configEntry: editedEntry });
      setGlobalError(err instanceof Error ? err.message : String(err));
    }
  };

  const valueType = watch('valueType');

  return (
    <Modal isOpen onClose={onClose}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <ModalOverlay />
        <ModalContent minW="2xl">
          <ModalHeader>{`Edit ${editedEntry.name}`}</ModalHeader>
          <ModalBody>
            <Text mb={6}>{editedEntry.documentation}</Text>

            <Flex gap={4} flexDirection="column">
              <FormField label={editedEntry.name}>
                <Controller
                  control={control}
                  name="valueType"
                  render={({ field: { onChange, value } }) => (
                    <RadioGroup<Inputs['valueType']> options={valueTypeOptions} value={value} onChange={onChange} />
                  )}
                />
              </FormField>
              {valueType === 'custom' && (
                <FormField label={`Set a custom ${editedEntry.name} value for this topic`}>
                  <Box maxW="fit-content">
                    <Controller
                      control={control}
                      name="customValue"
                      render={({ field: { onChange, value } }) => (
                        <ConfigEntryEditorController entry={editedEntry} onChange={onChange} value={value} />
                      )}
                    />
                  </Box>
                </FormField>
              )}
              {/*It's not possible to show default value until we get it always from the BE.*/}
              {/*Currently we only retrieve the current value and not default if it's set to custom/infinite*/}
              {/*{valueType === 'default' && (*/}
              {/*  <Box>*/}
              {/*    The default value is{' '}*/}
              {/*    <Text fontWeight="bold" display="inline">*/}
              {/*      {formatConfigValue(editedEntry.name, editedEntry.value, 'friendly')}*/}
              {/*    </Text>*/}
              {/*    . This is inherited from {editedEntry.source}.*/}
              {/*  </Box>*/}
              {/*)}*/}
            </Flex>
            {globalError && (
              <Alert status="error" my={2}>
                <AlertIcon />
                {globalError}
              </Alert>
            )}
          </ModalBody>
          <ModalFooter display="flex" gap={2}>
            <Button
              variant="ghost"
              onClick={() => {
                onClose();
              }}
            >
              Cancel
            </Button>
            <Button variant="solid" type="submit">
              Save changes
            </Button>
          </ModalFooter>
        </ModalContent>
      </form>
    </Modal>
  );
};

const ConfigurationEditor: FC<ConfigurationEditorProps> = observer((props) => {
  const $state = useLocalObservable<{
    filter?: string;
    editedEntry: ConfigEntryExtended | null;
  }>(() => ({
    filter: '',
    editedEntry: null,
  }));

  const editConfig = (configEntry: ConfigEntryExtended) => {
    $state.editedEntry = configEntry;
  };

  const topic = props.targetTopic;
  const hasEditPermissions = topic ? (api.topicPermissions.get(topic)?.canEditTopicConfig ?? true) : true;

  let entries = props.entries;
  const filter = $state.filter;
  if (filter) entries = entries.filter((x) => x.name.includes(filter) || (x.value ?? '').includes(filter));

  const entryOrder = {
    retention: -3,
    cleanup: -2,
  };

  entries = entries.slice().sort((a, b) => {
    for (const [e, order] of Object.entries(entryOrder)) {
      if (a.name.includes(e) && !b.name.includes(e)) return order;
      if (b.name.includes(e) && !a.name.includes(e)) return -order;
    }
    return 0;
  });

  const categories = entries.groupInto((x) => x.category);
  for (const e of categories) if (!e.key) e.key = 'Other';

  return (
    <Box pt={4}>
      {$state.editedEntry !== null && (
        <ConfigEditorForm
          targetTopic={props.targetTopic}
          editedEntry={$state.editedEntry}
          onClose={() => ($state.editedEntry = null)}
          onSuccess={() => {
            props.onForceRefresh();
          }}
        />
      )}
      <div className="configGroupTable">
        <SearchField
          searchText={$state.filter || ''}
          placeholderText="Filter"
          setSearchText={(value) => ($state.filter = value)}
          icon="filter"
        />
        {categories.map((x) => (
          <ConfigGroup
            key={x.key}
            groupName={x.key}
            entries={x.items}
            onEditEntry={editConfig}
            hasEditPermissions={hasEditPermissions}
          />
        ))}
      </div>
    </Box>
  );
});

export default ConfigurationEditor;

const ConfigGroup = observer(
  (p: {
    groupName?: string;
    onEditEntry: (configEntry: ConfigEntryExtended) => void;
    entries: ConfigEntryExtended[];
    hasEditPermissions: boolean;
  }) => {
    return (
      <>
        <div className="configGroupSpacer" />
        {p.groupName && <div className="configGroupTitle">{p.groupName}</div>}
        {p.entries.map((e) => (
          <ConfigEntry key={e.name} entry={e} onEditEntry={p.onEditEntry} hasEditPermissions={p.hasEditPermissions} />
        ))}
      </>
    );
  },
);

const ConfigEntry = observer(
  (p: {
    onEditEntry: (configEntry: ConfigEntryExtended) => void;
    entry: ConfigEntryExtended;
    hasEditPermissions: boolean;
  }) => {
    const { canEdit, reason: nonEdittableReason } = isTopicConfigEdittable(p.entry, p.hasEditPermissions);

    const entry = p.entry;
    const friendlyValue = formatConfigValue(entry.name, entry.value, 'friendly');

    return (
      <>
        <Flex direction="column">
          <Text fontWeight="600">{p.entry.name}</Text>
        </Flex>

        <Text>{friendlyValue}</Text>

        <span className="isEditted">{entry.isExplicitlySet && 'Custom'}</span>

        <span className="configButtons">
          <Tooltip label={nonEdittableReason} placement="left" isDisabled={canEdit} hasArrow>
            <span
              className={`btnEdit${canEdit ? '' : ' disabled'}`}
              onClick={() => {
                if (canEdit) p.onEditEntry(p.entry);
              }}
            >
              <Icon as={PencilIcon} />
            </span>
          </Tooltip>
          {entry.documentation && (
            <Popover
              hideCloseButton
              size="lg"
              content={
                <Flex flexDirection="column" gap={2}>
                  <Text fontSize="lg" fontWeight="bold">
                    {entry.name}
                  </Text>
                  <Text fontSize="sm">{entry.documentation}</Text>
                  <Text fontSize="sm">{getConfigDescription(entry.source)}</Text>
                </Flex>
              }
            >
              <Box>
                <Icon as={MdInfoOutline} />
              </Box>
            </Popover>
          )}
        </span>
      </>
    );
  },
);

function isTopicConfigEdittable(
  entry: ConfigEntryExtended,
  hasEditPermissions: boolean,
): { canEdit: boolean; reason?: string } {
  if (!hasEditPermissions)
    return { canEdit: false, reason: "You don't have permissions to change topic configuration entries" };

  if (isServerless()) {
    const edittableEntries = ['retention.ms', 'retention.bytes'];

    if (edittableEntries.includes(entry.name)) {
      return { canEdit: true };
    }

    return { canEdit: false, reason: 'This configuration is not editable on Serverless clusters' };
  }

  return { canEdit: true };
}

export const ConfigEntryEditorController = (p: {
  entry: ConfigEntryExtended;
  value: T;
  onChange: (e: T) => void;
}) => {
  const { entry, value, onChange } = p;
  switch (entry.frontendFormat) {
    case 'BOOLEAN':
      return (
        <SingleSelect
          options={[
            { value: 'false', label: 'False' },
            { value: 'true', label: 'True' },
          ]}
          value={value}
          onChange={onChange}
        />
      );

    case 'SELECT':
      return (
        <SingleSelect
          value={value}
          onChange={onChange}
          className={p.className}
          options={
            entry.enumValues?.map((value) => ({
              value,
              label: value,
            })) ?? []
          }
        />
      );

    case 'BYTE_SIZE':
      return (
        <DataSizeSelect
          allowInfinite={false}
          valueBytes={Number(value ?? 0)}
          onChange={(e) => onChange(Math.round(e))}
        />
      );
    case 'DURATION':
      return (
        <DurationSelect
          allowInfinite={false}
          valueMilliseconds={Number(value ?? 0)}
          onChange={(e) => onChange(Math.round(e))}
        />
      );

    case 'PASSWORD':
      return <PasswordInput value={value ?? ''} onChange={(x) => onChange(x.target.value)} />;

    case 'RATIO':
      return <RatioInput value={Number(value)} onChange={(x) => onChange(x)} />;

    case 'INTEGER':
      return <NumInput value={Number(value)} onChange={(e) => onChange(Math.round(e ?? 0))} />;

    case 'DECIMAL':
      return <NumInput value={Number(value)} onChange={(e) => onChange(e)} />;
    default:
      return <Input value={String(value)} onChange={(e) => onChange(e.target.value)} />;
  }
};

function getConfigDescription(source: string): string {
  switch (source) {
    case 'DEFAULT_CONFIG':
      return 'Inherited from DEFAULT_CONFIG';
    case 'DYNAMIC_TOPIC_CONFIG':
      return 'This is a custom setting for this topic';
    case 'DYNAMIC_BROKER_CONFIG':
    case 'STATIC_BROKER_CONFIG':
      return 'This is a custom setting set on the BROKER_CONFIG level.';
    default:
      return '';
  }
}
