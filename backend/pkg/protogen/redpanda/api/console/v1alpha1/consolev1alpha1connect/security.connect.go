// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: redpanda/api/console/v1alpha1/security.proto

package consolev1alpha1connect

import (
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"

	connect "connectrpc.com/connect"

	v1alpha1 "github.com/redpanda-data/console/backend/pkg/protogen/redpanda/api/console/v1alpha1"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// SecurityServiceName is the fully-qualified name of the SecurityService service.
	SecurityServiceName = "redpanda.api.console.v1alpha1.SecurityService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SecurityServiceListRolesProcedure is the fully-qualified name of the SecurityService's ListRoles
	// RPC.
	SecurityServiceListRolesProcedure = "/redpanda.api.console.v1alpha1.SecurityService/ListRoles"
	// SecurityServiceCreateRoleProcedure is the fully-qualified name of the SecurityService's
	// CreateRole RPC.
	SecurityServiceCreateRoleProcedure = "/redpanda.api.console.v1alpha1.SecurityService/CreateRole"
	// SecurityServiceGetRoleProcedure is the fully-qualified name of the SecurityService's GetRole RPC.
	SecurityServiceGetRoleProcedure = "/redpanda.api.console.v1alpha1.SecurityService/GetRole"
	// SecurityServiceUpdateRoleProcedure is the fully-qualified name of the SecurityService's
	// UpdateRole RPC.
	SecurityServiceUpdateRoleProcedure = "/redpanda.api.console.v1alpha1.SecurityService/UpdateRole"
	// SecurityServiceDeleteRoleProcedure is the fully-qualified name of the SecurityService's
	// DeleteRole RPC.
	SecurityServiceDeleteRoleProcedure = "/redpanda.api.console.v1alpha1.SecurityService/DeleteRole"
	// SecurityServiceListRoleMembersProcedure is the fully-qualified name of the SecurityService's
	// ListRoleMembers RPC.
	SecurityServiceListRoleMembersProcedure = "/redpanda.api.console.v1alpha1.SecurityService/ListRoleMembers"
	// SecurityServiceUpdateRoleMembershipProcedure is the fully-qualified name of the SecurityService's
	// UpdateRoleMembership RPC.
	SecurityServiceUpdateRoleMembershipProcedure = "/redpanda.api.console.v1alpha1.SecurityService/UpdateRoleMembership"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	securityServiceServiceDescriptor                    = v1alpha1.File_redpanda_api_console_v1alpha1_security_proto.Services().ByName("SecurityService")
	securityServiceListRolesMethodDescriptor            = securityServiceServiceDescriptor.Methods().ByName("ListRoles")
	securityServiceCreateRoleMethodDescriptor           = securityServiceServiceDescriptor.Methods().ByName("CreateRole")
	securityServiceGetRoleMethodDescriptor              = securityServiceServiceDescriptor.Methods().ByName("GetRole")
	securityServiceUpdateRoleMethodDescriptor           = securityServiceServiceDescriptor.Methods().ByName("UpdateRole")
	securityServiceDeleteRoleMethodDescriptor           = securityServiceServiceDescriptor.Methods().ByName("DeleteRole")
	securityServiceListRoleMembersMethodDescriptor      = securityServiceServiceDescriptor.Methods().ByName("ListRoleMembers")
	securityServiceUpdateRoleMembershipMethodDescriptor = securityServiceServiceDescriptor.Methods().ByName("UpdateRoleMembership")
)

// SecurityServiceClient is a client for the redpanda.api.console.v1alpha1.SecurityService service.
type SecurityServiceClient interface {
	// ListRoles lists all the roles based on optional filter.
	ListRoles(context.Context, *connect.Request[v1alpha1.ListRolesRequest]) (*connect.Response[v1alpha1.ListRolesResponse], error)
	CreateRole(context.Context, *connect.Request[v1alpha1.CreateRoleRequest]) (*connect.Response[v1alpha1.CreateRoleResponse], error)
	// GetRole retrieves the specific role.
	GetRole(context.Context, *connect.Request[v1alpha1.GetRoleRequest]) (*connect.Response[v1alpha1.GetRoleResponse], error)
	UpdateRole(context.Context, *connect.Request[v1alpha1.UpdateRoleRequest]) (*connect.Response[v1alpha1.UpdateRoleResponse], error)
	// DeleteRole deletes the role from the system.
	DeleteRole(context.Context, *connect.Request[v1alpha1.DeleteRoleRequest]) (*connect.Response[v1alpha1.DeleteRoleResponse], error)
	// ListRoleMembership lists all the members assigned to a role based on optional filter.
	ListRoleMembers(context.Context, *connect.Request[v1alpha1.ListRoleMembersRequest]) (*connect.Response[v1alpha1.ListRoleMembersResponse], error)
	// UpdateRoleMembership updates role membership.
	// Partially update role membership, adding or removing from a role
	// ONLY those members listed in the “add” or “remove” fields, respectively.
	// Adding a member that is already assigned to the role (or removing one that is not) is a no-op,
	// and the rest of the members will be added and removed and reported.
	UpdateRoleMembership(context.Context, *connect.Request[v1alpha1.UpdateRoleMembershipRequest]) (*connect.Response[v1alpha1.UpdateRoleMembershipResponse], error)
}

// NewSecurityServiceClient constructs a client for the
// redpanda.api.console.v1alpha1.SecurityService service. By default, it uses the Connect protocol
// with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To
// use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb()
// options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSecurityServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SecurityServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &securityServiceClient{
		listRoles: connect.NewClient[v1alpha1.ListRolesRequest, v1alpha1.ListRolesResponse](
			httpClient,
			baseURL+SecurityServiceListRolesProcedure,
			connect.WithSchema(securityServiceListRolesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createRole: connect.NewClient[v1alpha1.CreateRoleRequest, v1alpha1.CreateRoleResponse](
			httpClient,
			baseURL+SecurityServiceCreateRoleProcedure,
			connect.WithSchema(securityServiceCreateRoleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getRole: connect.NewClient[v1alpha1.GetRoleRequest, v1alpha1.GetRoleResponse](
			httpClient,
			baseURL+SecurityServiceGetRoleProcedure,
			connect.WithSchema(securityServiceGetRoleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateRole: connect.NewClient[v1alpha1.UpdateRoleRequest, v1alpha1.UpdateRoleResponse](
			httpClient,
			baseURL+SecurityServiceUpdateRoleProcedure,
			connect.WithSchema(securityServiceUpdateRoleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteRole: connect.NewClient[v1alpha1.DeleteRoleRequest, v1alpha1.DeleteRoleResponse](
			httpClient,
			baseURL+SecurityServiceDeleteRoleProcedure,
			connect.WithSchema(securityServiceDeleteRoleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listRoleMembers: connect.NewClient[v1alpha1.ListRoleMembersRequest, v1alpha1.ListRoleMembersResponse](
			httpClient,
			baseURL+SecurityServiceListRoleMembersProcedure,
			connect.WithSchema(securityServiceListRoleMembersMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateRoleMembership: connect.NewClient[v1alpha1.UpdateRoleMembershipRequest, v1alpha1.UpdateRoleMembershipResponse](
			httpClient,
			baseURL+SecurityServiceUpdateRoleMembershipProcedure,
			connect.WithSchema(securityServiceUpdateRoleMembershipMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// securityServiceClient implements SecurityServiceClient.
type securityServiceClient struct {
	listRoles            *connect.Client[v1alpha1.ListRolesRequest, v1alpha1.ListRolesResponse]
	createRole           *connect.Client[v1alpha1.CreateRoleRequest, v1alpha1.CreateRoleResponse]
	getRole              *connect.Client[v1alpha1.GetRoleRequest, v1alpha1.GetRoleResponse]
	updateRole           *connect.Client[v1alpha1.UpdateRoleRequest, v1alpha1.UpdateRoleResponse]
	deleteRole           *connect.Client[v1alpha1.DeleteRoleRequest, v1alpha1.DeleteRoleResponse]
	listRoleMembers      *connect.Client[v1alpha1.ListRoleMembersRequest, v1alpha1.ListRoleMembersResponse]
	updateRoleMembership *connect.Client[v1alpha1.UpdateRoleMembershipRequest, v1alpha1.UpdateRoleMembershipResponse]
}

// ListRoles calls redpanda.api.console.v1alpha1.SecurityService.ListRoles.
func (c *securityServiceClient) ListRoles(ctx context.Context, req *connect.Request[v1alpha1.ListRolesRequest]) (*connect.Response[v1alpha1.ListRolesResponse], error) {
	return c.listRoles.CallUnary(ctx, req)
}

// CreateRole calls redpanda.api.console.v1alpha1.SecurityService.CreateRole.
func (c *securityServiceClient) CreateRole(ctx context.Context, req *connect.Request[v1alpha1.CreateRoleRequest]) (*connect.Response[v1alpha1.CreateRoleResponse], error) {
	return c.createRole.CallUnary(ctx, req)
}

// GetRole calls redpanda.api.console.v1alpha1.SecurityService.GetRole.
func (c *securityServiceClient) GetRole(ctx context.Context, req *connect.Request[v1alpha1.GetRoleRequest]) (*connect.Response[v1alpha1.GetRoleResponse], error) {
	return c.getRole.CallUnary(ctx, req)
}

// UpdateRole calls redpanda.api.console.v1alpha1.SecurityService.UpdateRole.
func (c *securityServiceClient) UpdateRole(ctx context.Context, req *connect.Request[v1alpha1.UpdateRoleRequest]) (*connect.Response[v1alpha1.UpdateRoleResponse], error) {
	return c.updateRole.CallUnary(ctx, req)
}

// DeleteRole calls redpanda.api.console.v1alpha1.SecurityService.DeleteRole.
func (c *securityServiceClient) DeleteRole(ctx context.Context, req *connect.Request[v1alpha1.DeleteRoleRequest]) (*connect.Response[v1alpha1.DeleteRoleResponse], error) {
	return c.deleteRole.CallUnary(ctx, req)
}

// ListRoleMembers calls redpanda.api.console.v1alpha1.SecurityService.ListRoleMembers.
func (c *securityServiceClient) ListRoleMembers(ctx context.Context, req *connect.Request[v1alpha1.ListRoleMembersRequest]) (*connect.Response[v1alpha1.ListRoleMembersResponse], error) {
	return c.listRoleMembers.CallUnary(ctx, req)
}

// UpdateRoleMembership calls redpanda.api.console.v1alpha1.SecurityService.UpdateRoleMembership.
func (c *securityServiceClient) UpdateRoleMembership(ctx context.Context, req *connect.Request[v1alpha1.UpdateRoleMembershipRequest]) (*connect.Response[v1alpha1.UpdateRoleMembershipResponse], error) {
	return c.updateRoleMembership.CallUnary(ctx, req)
}

// SecurityServiceHandler is an implementation of the redpanda.api.console.v1alpha1.SecurityService
// service.
type SecurityServiceHandler interface {
	// ListRoles lists all the roles based on optional filter.
	ListRoles(context.Context, *connect.Request[v1alpha1.ListRolesRequest]) (*connect.Response[v1alpha1.ListRolesResponse], error)
	CreateRole(context.Context, *connect.Request[v1alpha1.CreateRoleRequest]) (*connect.Response[v1alpha1.CreateRoleResponse], error)
	// GetRole retrieves the specific role.
	GetRole(context.Context, *connect.Request[v1alpha1.GetRoleRequest]) (*connect.Response[v1alpha1.GetRoleResponse], error)
	UpdateRole(context.Context, *connect.Request[v1alpha1.UpdateRoleRequest]) (*connect.Response[v1alpha1.UpdateRoleResponse], error)
	// DeleteRole deletes the role from the system.
	DeleteRole(context.Context, *connect.Request[v1alpha1.DeleteRoleRequest]) (*connect.Response[v1alpha1.DeleteRoleResponse], error)
	// ListRoleMembership lists all the members assigned to a role based on optional filter.
	ListRoleMembers(context.Context, *connect.Request[v1alpha1.ListRoleMembersRequest]) (*connect.Response[v1alpha1.ListRoleMembersResponse], error)
	// UpdateRoleMembership updates role membership.
	// Partially update role membership, adding or removing from a role
	// ONLY those members listed in the “add” or “remove” fields, respectively.
	// Adding a member that is already assigned to the role (or removing one that is not) is a no-op,
	// and the rest of the members will be added and removed and reported.
	UpdateRoleMembership(context.Context, *connect.Request[v1alpha1.UpdateRoleMembershipRequest]) (*connect.Response[v1alpha1.UpdateRoleMembershipResponse], error)
}

// NewSecurityServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSecurityServiceHandler(svc SecurityServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	securityServiceListRolesHandler := connect.NewUnaryHandler(
		SecurityServiceListRolesProcedure,
		svc.ListRoles,
		connect.WithSchema(securityServiceListRolesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	securityServiceCreateRoleHandler := connect.NewUnaryHandler(
		SecurityServiceCreateRoleProcedure,
		svc.CreateRole,
		connect.WithSchema(securityServiceCreateRoleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	securityServiceGetRoleHandler := connect.NewUnaryHandler(
		SecurityServiceGetRoleProcedure,
		svc.GetRole,
		connect.WithSchema(securityServiceGetRoleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	securityServiceUpdateRoleHandler := connect.NewUnaryHandler(
		SecurityServiceUpdateRoleProcedure,
		svc.UpdateRole,
		connect.WithSchema(securityServiceUpdateRoleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	securityServiceDeleteRoleHandler := connect.NewUnaryHandler(
		SecurityServiceDeleteRoleProcedure,
		svc.DeleteRole,
		connect.WithSchema(securityServiceDeleteRoleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	securityServiceListRoleMembersHandler := connect.NewUnaryHandler(
		SecurityServiceListRoleMembersProcedure,
		svc.ListRoleMembers,
		connect.WithSchema(securityServiceListRoleMembersMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	securityServiceUpdateRoleMembershipHandler := connect.NewUnaryHandler(
		SecurityServiceUpdateRoleMembershipProcedure,
		svc.UpdateRoleMembership,
		connect.WithSchema(securityServiceUpdateRoleMembershipMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/redpanda.api.console.v1alpha1.SecurityService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SecurityServiceListRolesProcedure:
			securityServiceListRolesHandler.ServeHTTP(w, r)
		case SecurityServiceCreateRoleProcedure:
			securityServiceCreateRoleHandler.ServeHTTP(w, r)
		case SecurityServiceGetRoleProcedure:
			securityServiceGetRoleHandler.ServeHTTP(w, r)
		case SecurityServiceUpdateRoleProcedure:
			securityServiceUpdateRoleHandler.ServeHTTP(w, r)
		case SecurityServiceDeleteRoleProcedure:
			securityServiceDeleteRoleHandler.ServeHTTP(w, r)
		case SecurityServiceListRoleMembersProcedure:
			securityServiceListRoleMembersHandler.ServeHTTP(w, r)
		case SecurityServiceUpdateRoleMembershipProcedure:
			securityServiceUpdateRoleMembershipHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSecurityServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSecurityServiceHandler struct{}

func (UnimplementedSecurityServiceHandler) ListRoles(context.Context, *connect.Request[v1alpha1.ListRolesRequest]) (*connect.Response[v1alpha1.ListRolesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.console.v1alpha1.SecurityService.ListRoles is not implemented"))
}

func (UnimplementedSecurityServiceHandler) CreateRole(context.Context, *connect.Request[v1alpha1.CreateRoleRequest]) (*connect.Response[v1alpha1.CreateRoleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.console.v1alpha1.SecurityService.CreateRole is not implemented"))
}

func (UnimplementedSecurityServiceHandler) GetRole(context.Context, *connect.Request[v1alpha1.GetRoleRequest]) (*connect.Response[v1alpha1.GetRoleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.console.v1alpha1.SecurityService.GetRole is not implemented"))
}

func (UnimplementedSecurityServiceHandler) UpdateRole(context.Context, *connect.Request[v1alpha1.UpdateRoleRequest]) (*connect.Response[v1alpha1.UpdateRoleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.console.v1alpha1.SecurityService.UpdateRole is not implemented"))
}

func (UnimplementedSecurityServiceHandler) DeleteRole(context.Context, *connect.Request[v1alpha1.DeleteRoleRequest]) (*connect.Response[v1alpha1.DeleteRoleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.console.v1alpha1.SecurityService.DeleteRole is not implemented"))
}

func (UnimplementedSecurityServiceHandler) ListRoleMembers(context.Context, *connect.Request[v1alpha1.ListRoleMembersRequest]) (*connect.Response[v1alpha1.ListRoleMembersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.console.v1alpha1.SecurityService.ListRoleMembers is not implemented"))
}

func (UnimplementedSecurityServiceHandler) UpdateRoleMembership(context.Context, *connect.Request[v1alpha1.UpdateRoleMembershipRequest]) (*connect.Response[v1alpha1.UpdateRoleMembershipResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.console.v1alpha1.SecurityService.UpdateRoleMembership is not implemented"))
}