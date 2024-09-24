// @generated by protoc-gen-connect-es v1.2.0 with parameter "target=ts,import_extension="
// @generated from file redpanda/api/dataplane/v1alpha2/topic.proto (package redpanda.api.dataplane.v1alpha2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateTopicRequest, CreateTopicResponse, DeleteTopicRequest, DeleteTopicResponse, GetTopicConfigurationsRequest, GetTopicConfigurationsResponse, ListTopicsRequest, ListTopicsResponse, MountTopicsRequest, MountTopicsResponse, SetTopicConfigurationsRequest, SetTopicConfigurationsResponse, UnmountTopicsRequest, UnmountTopicsResponse, UpdateTopicConfigurationsRequest, UpdateTopicConfigurationsResponse } from "./topic_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service redpanda.api.dataplane.v1alpha2.TopicService
 */
export const TopicService = {
  typeName: "redpanda.api.dataplane.v1alpha2.TopicService",
  methods: {
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha2.TopicService.CreateTopic
     */
    createTopic: {
      name: "CreateTopic",
      I: CreateTopicRequest,
      O: CreateTopicResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha2.TopicService.ListTopics
     */
    listTopics: {
      name: "ListTopics",
      I: ListTopicsRequest,
      O: ListTopicsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha2.TopicService.DeleteTopic
     */
    deleteTopic: {
      name: "DeleteTopic",
      I: DeleteTopicRequest,
      O: DeleteTopicResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha2.TopicService.GetTopicConfigurations
     */
    getTopicConfigurations: {
      name: "GetTopicConfigurations",
      I: GetTopicConfigurationsRequest,
      O: GetTopicConfigurationsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha2.TopicService.UpdateTopicConfigurations
     */
    updateTopicConfigurations: {
      name: "UpdateTopicConfigurations",
      I: UpdateTopicConfigurationsRequest,
      O: UpdateTopicConfigurationsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha2.TopicService.SetTopicConfigurations
     */
    setTopicConfigurations: {
      name: "SetTopicConfigurations",
      I: SetTopicConfigurationsRequest,
      O: SetTopicConfigurationsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha2.TopicService.MountTopics
     */
    mountTopics: {
      name: "MountTopics",
      I: MountTopicsRequest,
      O: MountTopicsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha2.TopicService.UnmountTopics
     */
    unmountTopics: {
      name: "UnmountTopics",
      I: UnmountTopicsRequest,
      O: UnmountTopicsResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

