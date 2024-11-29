// @generated by protoc-gen-es v1.6.0 with parameter "target=ts,import_extension="
// @generated from file redpanda/api/dataplane/v1alpha2/error.proto (package redpanda.api.dataplane.v1alpha2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum redpanda.api.dataplane.v1alpha2.Reason
 */
export enum Reason {
	/**
	 * @generated from enum value: REASON_UNSPECIFIED = 0;
	 */
	UNSPECIFIED = 0,

	/**
	 * The feature is not configured.
	 *
	 * @generated from enum value: REASON_FEATURE_NOT_CONFIGURED = 1;
	 */
	FEATURE_NOT_CONFIGURED = 1,

	/**
	 * Internal Redpanda Console or data plane error.
	 *
	 * @generated from enum value: REASON_CONSOLE_ERROR = 2;
	 */
	CONSOLE_ERROR = 2,

	/**
	 * Redpanda Admin API returned an error.
	 *
	 * @generated from enum value: REASON_REDPANDA_ADMIN_API_ERROR = 3;
	 */
	REDPANDA_ADMIN_API_ERROR = 3,

	/**
	 * Redpanda or Kafka protocol error.
	 *
	 * @generated from enum value: REASON_KAFKA_API_ERROR = 4;
	 */
	KAFKA_API_ERROR = 4,

	/**
	 * Kafka Connect API error.
	 *
	 * @generated from enum value: REASON_KAFKA_CONNECT_API_ERROR = 5;
	 */
	KAFKA_CONNECT_API_ERROR = 5,

	/**
	 * Type mapping error translating internal or external types to API types.
	 *
	 * @generated from enum value: REASON_TYPE_MAPPING_ERROR = 6;
	 */
	TYPE_MAPPING_ERROR = 6,

	/**
	 * Cloud provider's secret store manager error.
	 *
	 * @generated from enum value: REASON_SECRET_STORE_ERROR = 7;
	 */
	SECRET_STORE_ERROR = 7,

	/**
	 * Invalid pipeline configuration.
	 *
	 * @generated from enum value: REASON_CONNECT_INVALID_PIPELINE_CONFIGURATION = 8;
	 */
	CONNECT_INVALID_PIPELINE_CONFIGURATION = 8,

	/**
	 * The Redpanda enterprise license has expired and is no longer valid.
	 *
	 * @generated from enum value: REASON_ENTERPRISE_LICENSE_EXPIRED = 9;
	 */
	ENTERPRISE_LICENSE_EXPIRED = 9,
}
// Retrieve enum metadata with: proto3.getEnumType(Reason)
proto3.util.setEnumType(Reason, "redpanda.api.dataplane.v1alpha2.Reason", [
	{ no: 0, name: "REASON_UNSPECIFIED" },
	{ no: 1, name: "REASON_FEATURE_NOT_CONFIGURED" },
	{ no: 2, name: "REASON_CONSOLE_ERROR" },
	{ no: 3, name: "REASON_REDPANDA_ADMIN_API_ERROR" },
	{ no: 4, name: "REASON_KAFKA_API_ERROR" },
	{ no: 5, name: "REASON_KAFKA_CONNECT_API_ERROR" },
	{ no: 6, name: "REASON_TYPE_MAPPING_ERROR" },
	{ no: 7, name: "REASON_SECRET_STORE_ERROR" },
	{ no: 8, name: "REASON_CONNECT_INVALID_PIPELINE_CONFIGURATION" },
	{ no: 9, name: "REASON_ENTERPRISE_LICENSE_EXPIRED" },
]);
