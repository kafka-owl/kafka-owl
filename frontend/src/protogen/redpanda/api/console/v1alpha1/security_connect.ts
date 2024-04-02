// @generated by protoc-gen-connect-es v1.2.0 with parameter "target=ts,import_extension="
// @generated from file redpanda/api/console/v1alpha1/security.proto (package redpanda.api.console.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateRoleRequest, CreateRoleResponse, DeleteRoleRequest, DeleteRoleResponse, GetRoleRequest, GetRoleResponse, ListRoleMembersRequest, ListRoleMembersResponse, ListRolesRequest, ListRolesResponse, ListRolesWithMembersRequest, ListRolesWithMembersResponse, UpdateRoleMembershipRequest, UpdateRoleMembershipResponse, UpdateRoleRequest, UpdateRoleResponse } from "./security_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service redpanda.api.console.v1alpha1.SecurityService
 */
export const SecurityService = {
  typeName: "redpanda.api.console.v1alpha1.SecurityService",
  methods: {
    /**
     * ListRoles lists all the roles based on optional filter.
     *
     * @generated from rpc redpanda.api.console.v1alpha1.SecurityService.ListRoles
     */
    listRoles: {
      name: "ListRoles",
      I: ListRolesRequest,
      O: ListRolesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.console.v1alpha1.SecurityService.CreateRole
     */
    createRole: {
      name: "CreateRole",
      I: CreateRoleRequest,
      O: CreateRoleResponse,
      kind: MethodKind.Unary,
    },
    /**
     * GetRole retrieves the specific role.
     *
     * @generated from rpc redpanda.api.console.v1alpha1.SecurityService.GetRole
     */
    getRole: {
      name: "GetRole",
      I: GetRoleRequest,
      O: GetRoleResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.console.v1alpha1.SecurityService.UpdateRole
     */
    updateRole: {
      name: "UpdateRole",
      I: UpdateRoleRequest,
      O: UpdateRoleResponse,
      kind: MethodKind.Unary,
    },
    /**
     * DeleteRole deletes the role from the system.
     *
     * @generated from rpc redpanda.api.console.v1alpha1.SecurityService.DeleteRole
     */
    deleteRole: {
      name: "DeleteRole",
      I: DeleteRoleRequest,
      O: DeleteRoleResponse,
      kind: MethodKind.Unary,
    },
    /**
     * ListRoleMembership lists all the members assigned to a role based on optional filter.
     *
     * @generated from rpc redpanda.api.console.v1alpha1.SecurityService.ListRoleMembers
     */
    listRoleMembers: {
      name: "ListRoleMembers",
      I: ListRoleMembersRequest,
      O: ListRoleMembersResponse,
      kind: MethodKind.Unary,
    },
    /**
     * UpdateRoleMembership updates role membership.
     * Partially update role membership, adding or removing from a role
     * ONLY those members listed in the “add” or “remove” fields, respectively.
     * Adding a member that is already assigned to the role (or removing one that is not) is a no-op,
     * and the rest of the members will be added and removed and reported.
     *
     * @generated from rpc redpanda.api.console.v1alpha1.SecurityService.UpdateRoleMembership
     */
    updateRoleMembership: {
      name: "UpdateRoleMembership",
      I: UpdateRoleMembershipRequest,
      O: UpdateRoleMembershipResponse,
      kind: MethodKind.Unary,
    },
    /**
     * ListRolesWithMembers lists all the roles and their members based on optional filter.
     *
     * @generated from rpc redpanda.api.console.v1alpha1.SecurityService.ListRolesWithMembers
     */
    listRolesWithMembers: {
      name: "ListRolesWithMembers",
      I: ListRolesWithMembersRequest,
      O: ListRolesWithMembersResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

