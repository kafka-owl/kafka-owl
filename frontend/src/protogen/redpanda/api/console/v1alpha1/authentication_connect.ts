// @generated by protoc-gen-connect-es v1.2.0 with parameter "target=ts,import_extension="
// @generated from file redpanda/api/console/v1alpha1/authentication.proto (package redpanda.api.console.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { ListAuthenticationMethodsRequest, ListAuthenticationMethodsResponse, LoginSaslScramRequest, LoginSaslScramResponse } from "./authentication_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service redpanda.api.console.v1alpha1.AuthenticationService
 */
export const AuthenticationService = {
  typeName: "redpanda.api.console.v1alpha1.AuthenticationService",
  methods: {
    /**
     * RPC to login using SASL-SCRAM.
     *
     * @generated from rpc redpanda.api.console.v1alpha1.AuthenticationService.LoginSaslScram
     */
    loginSaslScram: {
      name: "LoginSaslScram",
      I: LoginSaslScramRequest,
      O: LoginSaslScramResponse,
      kind: MethodKind.Unary,
    },
    /**
     * RPC to list available authentication methods.
     *
     * @generated from rpc redpanda.api.console.v1alpha1.AuthenticationService.ListAuthenticationMethods
     */
    listAuthenticationMethods: {
      name: "ListAuthenticationMethods",
      I: ListAuthenticationMethodsRequest,
      O: ListAuthenticationMethodsResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

