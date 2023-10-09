// @generated by protoc-gen-es v1.3.3 with parameter "target=ts,import_extension="
// @generated from file redpanda/api/common/v1alpha1/common.proto (package redpanda.api.common.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Any, Message, proto3 } from "@bufbuild/protobuf";
import { Code } from "../../../../google/rpc/code_pb";

/**
 * @generated from enum redpanda.api.common.v1alpha1.Reason
 */
export enum Reason {
  /**
   * @generated from enum value: REASON_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,
}
// Retrieve enum metadata with: proto3.getEnumType(Reason)
proto3.util.setEnumType(Reason, "redpanda.api.common.v1alpha1.Reason", [
  { no: 0, name: "REASON_UNSPECIFIED" },
]);

/**
 * Modified variant of google.rpc.Status, that uses enum instead of int32 for
 * code, so it's nicer in REST.
 * The `Status` type defines a logical error model that is suitable for
 * different programming environments, including REST APIs and RPC APIs. It is
 * used by [gRPC](https://github.com/grpc). Each `Status` message contains
 * three pieces of data: error code, error message, and error details.
 *
 * You can find out more about this error model and how to work with it in the
 * [API Design Guide](https://cloud.google.com/apis/design/errors).
 *
 * @generated from message redpanda.api.common.v1alpha1.ErrorStatus
 */
export class ErrorStatus extends Message<ErrorStatus> {
  /**
   * The status code, which should be an enum value of
   * [google.rpc.Code][google.rpc.Code].
   *
   * @generated from field: google.rpc.Code code = 1;
   */
  code = Code.OK;

  /**
   * A developer-facing error message, which should be in English. Any
   * user-facing error message should be localized and sent in the
   * [google.rpc.Status.details][google.rpc.Status.details] field, or localized
   * by the client.
   *
   * @generated from field: string message = 2;
   */
  message = "";

  /**
   * A list of messages that carry the error details.  There is a common set of
   * message types for APIs to use.
   *
   * @generated from field: repeated google.protobuf.Any details = 3;
   */
  details: Any[] = [];

  constructor(data?: PartialMessage<ErrorStatus>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.common.v1alpha1.ErrorStatus";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "code", kind: "enum", T: proto3.getEnumType(Code) },
    { no: 2, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "details", kind: "message", T: Any, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ErrorStatus {
    return new ErrorStatus().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ErrorStatus {
    return new ErrorStatus().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ErrorStatus {
    return new ErrorStatus().fromJsonString(jsonString, options);
  }

  static equals(a: ErrorStatus | PlainMessage<ErrorStatus> | undefined, b: ErrorStatus | PlainMessage<ErrorStatus> | undefined): boolean {
    return proto3.util.equals(ErrorStatus, a, b);
  }
}

