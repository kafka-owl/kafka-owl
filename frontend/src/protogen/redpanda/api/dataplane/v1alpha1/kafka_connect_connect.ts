// @generated by protoc-gen-connect-es v1.2.0 with parameter "target=ts,import_extension="
// @generated from file redpanda/api/dataplane/v1alpha1/kafka_connect.proto (package redpanda.api.dataplane.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateConnectorRequest, CreateConnectorResponse, DeleteConnectorRequest, GetConnectClusterRequest, GetConnectClusterResponse, GetConnectorConfigRequest, GetConnectorConfigResponse, GetConnectorRequest, GetConnectorResponse, ListConnectClustersRequest, ListConnectClustersResponse, ListConnectorsRequest, ListConnectorsResponse, ListConnectorTopicsRequest, ListConnectorTopicsResponse, PauseConnectorRequest, ResetConnectorTopicsRequest, RestartConnectorRequest, ResumeConnectorRequest, UpsertConnectorRequest, UpsertConnectorResponse } from "./kafka_connect_pb";
import { Empty, MethodKind } from "@bufbuild/protobuf";

/**
 * KafkaConnectService is the service for the Kafka connect, it exposes the
 * Kafka Connect API, you can set multiple Kafka connect services and all of
 * them can be managed using this service definition, the request is not only
 * proxied but also enriched with better error handling and custom
 * documentation and configuration
 *
 * @generated from service redpanda.api.dataplane.v1alpha1.KafkaConnectService
 */
export const KafkaConnectService = {
  typeName: "redpanda.api.dataplane.v1alpha1.KafkaConnectService",
  methods: {
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha1.KafkaConnectService.ListConnectClusters
     */
    listConnectClusters: {
      name: "ListConnectClusters",
      I: ListConnectClustersRequest,
      O: ListConnectClustersResponse,
      kind: MethodKind.Unary,
    },
    /**
     * GetConnectClusterInfo implements the get cluster info method, exposes a Kafka
     * Connect equivalent REST endpoint
     *
     * @generated from rpc redpanda.api.dataplane.v1alpha1.KafkaConnectService.GetConnectCluster
     */
    getConnectCluster: {
      name: "GetConnectCluster",
      I: GetConnectClusterRequest,
      O: GetConnectClusterResponse,
      kind: MethodKind.Unary,
    },
    /**
     * ListConnectors implements the list connectors method, exposes a Kafka
     * Connect equivalent REST endpoint
     *
     * @generated from rpc redpanda.api.dataplane.v1alpha1.KafkaConnectService.ListConnectors
     */
    listConnectors: {
      name: "ListConnectors",
      I: ListConnectorsRequest,
      O: ListConnectorsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * CreateConnector implements the create connector method, and exposes an
     * equivalent REST endpoint as the Kafka connect API endpoint
     *
     * @generated from rpc redpanda.api.dataplane.v1alpha1.KafkaConnectService.CreateConnector
     */
    createConnector: {
      name: "CreateConnector",
      I: CreateConnectorRequest,
      O: CreateConnectorResponse,
      kind: MethodKind.Unary,
    },
    /**
     * RestartConnector implements the restart connector method, exposes a Kafka
     * Connect equivalent REST endpoint
     *
     * @generated from rpc redpanda.api.dataplane.v1alpha1.KafkaConnectService.RestartConnector
     */
    restartConnector: {
      name: "RestartConnector",
      I: RestartConnectorRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * GetConnector implements the get connector method, exposes a Kafka
     * Connect equivalent REST endpoint
     *
     * @generated from rpc redpanda.api.dataplane.v1alpha1.KafkaConnectService.GetConnector
     */
    getConnector: {
      name: "GetConnector",
      I: GetConnectorRequest,
      O: GetConnectorResponse,
      kind: MethodKind.Unary,
    },
    /**
     * PauseConnector implements the pause connector method, exposes a Kafka
     * connect equivalent REST endpoint
     *
     * @generated from rpc redpanda.api.dataplane.v1alpha1.KafkaConnectService.PauseConnector
     */
    pauseConnector: {
      name: "PauseConnector",
      I: PauseConnectorRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * ResumeConnector implements the resume connector method, exposes a Kafka
     * connect equivalent REST endpoint
     *
     * @generated from rpc redpanda.api.dataplane.v1alpha1.KafkaConnectService.ResumeConnector
     */
    resumeConnector: {
      name: "ResumeConnector",
      I: ResumeConnectorRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * DeleteConnector implements the delete connector method, exposes a Kafka
     * connect equivalent REST endpoint
     *
     * @generated from rpc redpanda.api.dataplane.v1alpha1.KafkaConnectService.DeleteConnector
     */
    deleteConnector: {
      name: "DeleteConnector",
      I: DeleteConnectorRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * UpsertConector implements the update or create connector method, it
     * exposes a kafka connect equivalent REST endpoint
     *
     * @generated from rpc redpanda.api.dataplane.v1alpha1.KafkaConnectService.UpsertConnector
     */
    upsertConnector: {
      name: "UpsertConnector",
      I: UpsertConnectorRequest,
      O: UpsertConnectorResponse,
      kind: MethodKind.Unary,
    },
    /**
     * GetConnectorConfig implements the get connector config method, expose a kafka connect equivalent REST endpoint
     *
     * @generated from rpc redpanda.api.dataplane.v1alpha1.KafkaConnectService.GetConnectorConfig
     */
    getConnectorConfig: {
      name: "GetConnectorConfig",
      I: GetConnectorConfigRequest,
      O: GetConnectorConfigResponse,
      kind: MethodKind.Unary,
    },
    /**
     * ListConnectorTopics implements the list connector topics method, expose a kafka connect equivalent REST endpoint
     *
     * @generated from rpc redpanda.api.dataplane.v1alpha1.KafkaConnectService.ListConnectorTopics
     */
    listConnectorTopics: {
      name: "ListConnectorTopics",
      I: ListConnectorTopicsRequest,
      O: ListConnectorTopicsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * ResetConnectorTopics implements the reset connector topics method, expose a kafka connect equivalent REST endpoint
     * the request body is empty.
     *
     * @generated from rpc redpanda.api.dataplane.v1alpha1.KafkaConnectService.ResetConnectorTopics
     */
    resetConnectorTopics: {
      name: "ResetConnectorTopics",
      I: ResetConnectorTopicsRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
  }
} as const;

