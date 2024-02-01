// @generated by protoc-gen-es v1.6.0 with parameter "target=ts,import_extension="
// @generated from file redpanda/api/common/v1alpha1/money.proto (package redpanda.api.common.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum redpanda.api.common.v1alpha1.Currency
 */
export enum Currency {
  /**
   * @generated from enum value: CURRENCY_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: CURRENCY_CURRENCY_USD_CENTS = 1;
   */
  CURRENCY_USD_CENTS = 1,

  /**
   * @generated from enum value: CURRENCY_CURRENCY_CRT = 2;
   */
  CURRENCY_CRT = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(Currency)
proto3.util.setEnumType(Currency, "redpanda.api.common.v1alpha1.Currency", [
  { no: 0, name: "CURRENCY_UNSPECIFIED" },
  { no: 1, name: "CURRENCY_CURRENCY_USD_CENTS" },
  { no: 2, name: "CURRENCY_CURRENCY_CRT" },
]);

/**
 * @generated from message redpanda.api.common.v1alpha1.Money
 */
export class Money extends Message<Money> {
  /**
   * Amount is a decimal number.
   * Examples:
   * 10
   * 10.15
   *
   * @generated from field: string amount = 1;
   */
  amount = "";

  /**
   * @generated from field: redpanda.api.common.v1alpha1.Currency currency = 2;
   */
  currency = Currency.UNSPECIFIED;

  constructor(data?: PartialMessage<Money>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.common.v1alpha1.Money";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "amount", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "currency", kind: "enum", T: proto3.getEnumType(Currency) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Money {
    return new Money().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Money {
    return new Money().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Money {
    return new Money().fromJsonString(jsonString, options);
  }

  static equals(a: Money | PlainMessage<Money> | undefined, b: Money | PlainMessage<Money> | undefined): boolean {
    return proto3.util.equals(Money, a, b);
  }
}

