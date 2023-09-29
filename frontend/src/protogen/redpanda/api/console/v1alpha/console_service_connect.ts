// @generated by protoc-gen-connect-es v1.0.0 with parameter "target=ts,import_extension="
// @generated from file redpanda/api/console/v1alpha/console_service.proto (package redpanda.api.console.v1alpha, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { ListMessagesRequest, ListMessagesResponse } from "./list_messages_pb";
import { MethodKind } from "@bufbuild/protobuf";
import { PublishMessageRequest, PublishMessageResponse } from "./publish_messages_pb";

/**
 * ConsoleService represents the Console API service.
 *
 * @generated from service redpanda.api.console.v1alpha.ConsoleService
 */
export const ConsoleService = {
  typeName: "redpanda.api.console.v1alpha.ConsoleService",
  methods: {
    /**
     * ListMessages lists the messages according to the requested query.
     *
     * @generated from rpc redpanda.api.console.v1alpha.ConsoleService.ListMessages
     */
    listMessages: {
      name: "ListMessages",
      I: ListMessagesRequest,
      O: ListMessagesResponse,
      kind: MethodKind.ServerStreaming,
    },
    /**
     * PublishMessage publishes message.
     *
     * @generated from rpc redpanda.api.console.v1alpha.ConsoleService.PublishMessage
     */
    publishMessage: {
      name: "PublishMessage",
      I: PublishMessageRequest,
      O: PublishMessageResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
