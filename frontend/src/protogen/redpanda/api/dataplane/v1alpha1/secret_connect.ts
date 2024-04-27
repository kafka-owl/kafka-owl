// @generated by protoc-gen-connect-es v1.2.0 with parameter "target=ts,import_extension="
// @generated from file redpanda/api/dataplane/v1alpha1/secret.proto (package redpanda.api.dataplane.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateConnectSecretRequest, CreateConnectSecretResponse, DeleteConnectSecretRequest, DeleteConnectSecretResponse, GetConnectSecretRequest, GetConnectSecretResponse, ListConnectSecretsRequest, ListConnectSecretsResponse, UpdateConnectSecretRequest, UpdateConnectSecretResponse } from "./secret_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service redpanda.api.dataplane.v1alpha1.SecretService
 */
export const SecretService = {
  typeName: "redpanda.api.dataplane.v1alpha1.SecretService",
  methods: {
    /**
     * GetConnectSecret retrieves the specific secret for a specific Connect.
     *
     * @generated from rpc redpanda.api.dataplane.v1alpha1.SecretService.GetConnectSecret
     */
    getConnectSecret: {
      name: "GetConnectSecret",
      I: GetConnectSecretRequest,
      O: GetConnectSecretResponse,
      kind: MethodKind.Unary,
    },
    /**
     * ListConnectSecrets lists the Connect secrets based on optional filter.
     *
     * @generated from rpc redpanda.api.dataplane.v1alpha1.SecretService.ListConnectSecrets
     */
    listConnectSecrets: {
      name: "ListConnectSecrets",
      I: ListConnectSecretsRequest,
      O: ListConnectSecretsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * CreateConnectSecret creates the secret for a Connect.
     *
     * @generated from rpc redpanda.api.dataplane.v1alpha1.SecretService.CreateConnectSecret
     */
    createConnectSecret: {
      name: "CreateConnectSecret",
      I: CreateConnectSecretRequest,
      O: CreateConnectSecretResponse,
      kind: MethodKind.Unary,
    },
    /**
     * UpdateConnectSecret updates the Connect secret.
     *
     * @generated from rpc redpanda.api.dataplane.v1alpha1.SecretService.UpdateConnectSecret
     */
    updateConnectSecret: {
      name: "UpdateConnectSecret",
      I: UpdateConnectSecretRequest,
      O: UpdateConnectSecretResponse,
      kind: MethodKind.Unary,
    },
    /**
     * DeleteSecret deletes the secret.
     *
     * @generated from rpc redpanda.api.dataplane.v1alpha1.SecretService.DeleteConnectSecret
     */
    deleteConnectSecret: {
      name: "DeleteConnectSecret",
      I: DeleteConnectSecretRequest,
      O: DeleteConnectSecretResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

