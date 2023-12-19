// @generated by protoc-gen-connect-es v1.2.0 with parameter "target=ts,import_extension="
// @generated from file redpanda/api/dataplane/v1alpha1/acl.proto (package redpanda.api.dataplane.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateACLRequest, CreateACLResponse, DeleteACLsRequest, DeleteACLsResponse, ListACLsRequest, ListACLsResponse } from "./acl_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service redpanda.api.dataplane.v1alpha1.ACLService
 */
export const ACLService = {
  typeName: "redpanda.api.dataplane.v1alpha1.ACLService",
  methods: {
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha1.ACLService.ListACLs
     */
    listACLs: {
      name: "ListACLs",
      I: ListACLsRequest,
      O: ListACLsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha1.ACLService.CreateACL
     */
    createACL: {
      name: "CreateACL",
      I: CreateACLRequest,
      O: CreateACLResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha1.ACLService.DeleteACLs
     */
    deleteACLs: {
      name: "DeleteACLs",
      I: DeleteACLsRequest,
      O: DeleteACLsResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

