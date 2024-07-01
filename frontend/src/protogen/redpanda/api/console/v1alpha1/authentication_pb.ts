// @generated by protoc-gen-es v1.6.0 with parameter "target=ts,import_extension="
// @generated from file redpanda/api/console/v1alpha1/authentication.proto (package redpanda.api.console.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum redpanda.api.console.v1alpha1.SASLMechanism
 */
export enum SASLMechanism {
  /**
   * The SASL mechanism is unspecified.
   *
   * @generated from enum value: SASL_MECHANISM_UNSPECIFIED = 0;
   */
  SASL_MECHANISM_UNSPECIFIED = 0,

  /**
   * The SASL mechanism using SCRAM-SHA-256.
   *
   * @generated from enum value: SASL_MECHANISM_SCRAM_SHA_256 = 1;
   */
  SASL_MECHANISM_SCRAM_SHA_256 = 1,

  /**
   * The SASL mechanism using SCRAM-SHA-512.
   *
   * @generated from enum value: SASL_MECHANISM_SCRAM_SHA_512 = 2;
   */
  SASL_MECHANISM_SCRAM_SHA_512 = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(SASLMechanism)
proto3.util.setEnumType(SASLMechanism, "redpanda.api.console.v1alpha1.SASLMechanism", [
  { no: 0, name: "SASL_MECHANISM_UNSPECIFIED" },
  { no: 1, name: "SASL_MECHANISM_SCRAM_SHA_256" },
  { no: 2, name: "SASL_MECHANISM_SCRAM_SHA_512" },
]);

/**
 * @generated from enum redpanda.api.console.v1alpha1.AuthenticationMethod
 */
export enum AuthenticationMethod {
  /**
   * The authentication method is unspecified.
   *
   * @generated from enum value: AUTHENTICATION_METHOD_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * The authentication method using OpenID Connect.
   *
   * @generated from enum value: AUTHENTICATION_METHOD_OIDC = 1;
   */
  OIDC = 1,

  /**
   * The authentication method using SASL-SCRAM.
   *
   * @generated from enum value: AUTHENTICATION_METHOD_SASL_SCRAM = 2;
   */
  SASL_SCRAM = 2,

  /**
   * The authentication method for Redpanda Cloud.
   *
   * @generated from enum value: AUTHENTICATION_METHOD_REDPANDA_CLOUD = 3;
   */
  REDPANDA_CLOUD = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(AuthenticationMethod)
proto3.util.setEnumType(AuthenticationMethod, "redpanda.api.console.v1alpha1.AuthenticationMethod", [
  { no: 0, name: "AUTHENTICATION_METHOD_UNSPECIFIED" },
  { no: 1, name: "AUTHENTICATION_METHOD_OIDC" },
  { no: 2, name: "AUTHENTICATION_METHOD_SASL_SCRAM" },
  { no: 3, name: "AUTHENTICATION_METHOD_REDPANDA_CLOUD" },
]);

/**
 * @generated from message redpanda.api.console.v1alpha1.LoginSaslScramRequest
 */
export class LoginSaslScramRequest extends Message<LoginSaslScramRequest> {
  /**
   * The username for the login request.
   *
   * @generated from field: string username = 1;
   */
  username = "";

  /**
   * The password for the login request.
   *
   * @generated from field: string password = 2;
   */
  password = "";

  /**
   * The SASL mechanism to be used for authentication.
   *
   * @generated from field: redpanda.api.console.v1alpha1.SASLMechanism mechanism = 3;
   */
  mechanism = SASLMechanism.SASL_MECHANISM_UNSPECIFIED;

  constructor(data?: PartialMessage<LoginSaslScramRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.console.v1alpha1.LoginSaslScramRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "mechanism", kind: "enum", T: proto3.getEnumType(SASLMechanism) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LoginSaslScramRequest {
    return new LoginSaslScramRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LoginSaslScramRequest {
    return new LoginSaslScramRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LoginSaslScramRequest {
    return new LoginSaslScramRequest().fromJsonString(jsonString, options);
  }

  static equals(a: LoginSaslScramRequest | PlainMessage<LoginSaslScramRequest> | undefined, b: LoginSaslScramRequest | PlainMessage<LoginSaslScramRequest> | undefined): boolean {
    return proto3.util.equals(LoginSaslScramRequest, a, b);
  }
}

/**
 * @generated from message redpanda.api.console.v1alpha1.LoginSaslScramResponse
 */
export class LoginSaslScramResponse extends Message<LoginSaslScramResponse> {
  constructor(data?: PartialMessage<LoginSaslScramResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.console.v1alpha1.LoginSaslScramResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LoginSaslScramResponse {
    return new LoginSaslScramResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LoginSaslScramResponse {
    return new LoginSaslScramResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LoginSaslScramResponse {
    return new LoginSaslScramResponse().fromJsonString(jsonString, options);
  }

  static equals(a: LoginSaslScramResponse | PlainMessage<LoginSaslScramResponse> | undefined, b: LoginSaslScramResponse | PlainMessage<LoginSaslScramResponse> | undefined): boolean {
    return proto3.util.equals(LoginSaslScramResponse, a, b);
  }
}

/**
 * @generated from message redpanda.api.console.v1alpha1.ListAuthenticationMethodsRequest
 */
export class ListAuthenticationMethodsRequest extends Message<ListAuthenticationMethodsRequest> {
  constructor(data?: PartialMessage<ListAuthenticationMethodsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.console.v1alpha1.ListAuthenticationMethodsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListAuthenticationMethodsRequest {
    return new ListAuthenticationMethodsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListAuthenticationMethodsRequest {
    return new ListAuthenticationMethodsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListAuthenticationMethodsRequest {
    return new ListAuthenticationMethodsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListAuthenticationMethodsRequest | PlainMessage<ListAuthenticationMethodsRequest> | undefined, b: ListAuthenticationMethodsRequest | PlainMessage<ListAuthenticationMethodsRequest> | undefined): boolean {
    return proto3.util.equals(ListAuthenticationMethodsRequest, a, b);
  }
}

/**
 * @generated from message redpanda.api.console.v1alpha1.ListAuthenticationMethodsResponse
 */
export class ListAuthenticationMethodsResponse extends Message<ListAuthenticationMethodsResponse> {
  /**
   * The list of available authentication methods.
   *
   * @generated from field: repeated redpanda.api.console.v1alpha1.AuthenticationMethod methods = 1;
   */
  methods: AuthenticationMethod[] = [];

  constructor(data?: PartialMessage<ListAuthenticationMethodsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.console.v1alpha1.ListAuthenticationMethodsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "methods", kind: "enum", T: proto3.getEnumType(AuthenticationMethod), repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListAuthenticationMethodsResponse {
    return new ListAuthenticationMethodsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListAuthenticationMethodsResponse {
    return new ListAuthenticationMethodsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListAuthenticationMethodsResponse {
    return new ListAuthenticationMethodsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListAuthenticationMethodsResponse | PlainMessage<ListAuthenticationMethodsResponse> | undefined, b: ListAuthenticationMethodsResponse | PlainMessage<ListAuthenticationMethodsResponse> | undefined): boolean {
    return proto3.util.equals(ListAuthenticationMethodsResponse, a, b);
  }
}

