// @generated by protoc-gen-connect-es v1.2.0 with parameter "target=ts,import_extension="
// @generated from file redpanda/api/dataplane/v1alpha2/pipeline.proto (package redpanda.api.dataplane.v1alpha2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import {
	CreatePipelineRequest,
	CreatePipelineResponse,
	DeletePipelineRequest,
	DeletePipelineResponse,
	GetPipelineRequest,
	GetPipelineResponse,
	GetPipelineServiceConfigSchemaRequest,
	GetPipelineServiceConfigSchemaResponse,
	ListPipelinesRequest,
	ListPipelinesResponse,
	StartPipelineRequest,
	StartPipelineResponse,
	StopPipelineRequest,
	StopPipelineResponse,
	UpdatePipelineRequest,
	UpdatePipelineResponse,
} from "./pipeline_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * PipelineService is the service for Redpanda Connect.
 * It exposes the API for creating and managing Redpanda Connect pipelines and their configurations.
 *
 * @generated from service redpanda.api.dataplane.v1alpha2.PipelineService
 */
export const PipelineService = {
	typeName: "redpanda.api.dataplane.v1alpha2.PipelineService",
	methods: {
		/**
		 * CreatePipeline creates a Redpanda Connect pipeline in the Redpanda cluster.
		 *
		 * @generated from rpc redpanda.api.dataplane.v1alpha2.PipelineService.CreatePipeline
		 */
		createPipeline: {
			name: "CreatePipeline",
			I: CreatePipelineRequest,
			O: CreatePipelineResponse,
			kind: MethodKind.Unary,
		},
		/**
		 * GetPipeline gets a specific Redpanda Connect pipeline.
		 *
		 * @generated from rpc redpanda.api.dataplane.v1alpha2.PipelineService.GetPipeline
		 */
		getPipeline: {
			name: "GetPipeline",
			I: GetPipelineRequest,
			O: GetPipelineResponse,
			kind: MethodKind.Unary,
		},
		/**
		 * ListPipelines implements the list pipelines method which lists the pipelines
		 * in the Redpanda cluster.
		 *
		 * @generated from rpc redpanda.api.dataplane.v1alpha2.PipelineService.ListPipelines
		 */
		listPipelines: {
			name: "ListPipelines",
			I: ListPipelinesRequest,
			O: ListPipelinesResponse,
			kind: MethodKind.Unary,
		},
		/**
		 * UpdatePipeline updates a specific Redpanda Connect pipeline configuration.
		 *
		 * @generated from rpc redpanda.api.dataplane.v1alpha2.PipelineService.UpdatePipeline
		 */
		updatePipeline: {
			name: "UpdatePipeline",
			I: UpdatePipelineRequest,
			O: UpdatePipelineResponse,
			kind: MethodKind.Unary,
		},
		/**
		 * DeletePipeline deletes a specific Redpanda Connect pipeline.
		 *
		 * @generated from rpc redpanda.api.dataplane.v1alpha2.PipelineService.DeletePipeline
		 */
		deletePipeline: {
			name: "DeletePipeline",
			I: DeletePipelineRequest,
			O: DeletePipelineResponse,
			kind: MethodKind.Unary,
		},
		/**
		 * StopPipeline stops a specific Redpanda Connect pipeline.
		 *
		 * @generated from rpc redpanda.api.dataplane.v1alpha2.PipelineService.StopPipeline
		 */
		stopPipeline: {
			name: "StopPipeline",
			I: StopPipelineRequest,
			O: StopPipelineResponse,
			kind: MethodKind.Unary,
		},
		/**
		 * StartPipeline starts a specific Redpanda Connect pipeline that has been previously stopped.
		 *
		 * @generated from rpc redpanda.api.dataplane.v1alpha2.PipelineService.StartPipeline
		 */
		startPipeline: {
			name: "StartPipeline",
			I: StartPipelineRequest,
			O: StartPipelineResponse,
			kind: MethodKind.Unary,
		},
		/**
		 * The configuration schema includes available [components and processors](https://docs.redpanda.com/redpanda-cloud/develop/connect/components/about) in this Redpanda Connect instance.
		 *
		 * @generated from rpc redpanda.api.dataplane.v1alpha2.PipelineService.GetPipelineServiceConfigSchema
		 */
		getPipelineServiceConfigSchema: {
			name: "GetPipelineServiceConfigSchema",
			I: GetPipelineServiceConfigSchemaRequest,
			O: GetPipelineServiceConfigSchemaResponse,
			kind: MethodKind.Unary,
		},
	},
} as const;
