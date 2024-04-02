// Code generated by protoc-gen-connect-gateway. DO NOT EDIT.
//
// Source: redpanda/api/console/v1alpha1/security.proto

package consolev1alpha1connect

import (
	context "context"
	fmt "fmt"

	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	connect_gateway "go.vallahaye.net/connect-gateway"

	v1alpha1 "github.com/redpanda-data/console/backend/pkg/protogen/redpanda/api/console/v1alpha1"
)

// SecurityServiceGatewayServer implements the gRPC server API for the SecurityService service.
type SecurityServiceGatewayServer struct {
	v1alpha1.UnimplementedSecurityServiceServer
	listRoles            connect_gateway.UnaryHandler[v1alpha1.ListRolesRequest, v1alpha1.ListRolesResponse]
	createRole           connect_gateway.UnaryHandler[v1alpha1.CreateRoleRequest, v1alpha1.CreateRoleResponse]
	getRole              connect_gateway.UnaryHandler[v1alpha1.GetRoleRequest, v1alpha1.GetRoleResponse]
	updateRole           connect_gateway.UnaryHandler[v1alpha1.UpdateRoleRequest, v1alpha1.UpdateRoleResponse]
	deleteRole           connect_gateway.UnaryHandler[v1alpha1.DeleteRoleRequest, v1alpha1.DeleteRoleResponse]
	listRoleMembers      connect_gateway.UnaryHandler[v1alpha1.ListRoleMembersRequest, v1alpha1.ListRoleMembersResponse]
	updateRoleMembership connect_gateway.UnaryHandler[v1alpha1.UpdateRoleMembershipRequest, v1alpha1.UpdateRoleMembershipResponse]
	listRolesWithMembers connect_gateway.UnaryHandler[v1alpha1.ListRolesWithMembersRequest, v1alpha1.ListRolesWithMembersResponse]
}

// NewSecurityServiceGatewayServer constructs a Connect-Gateway gRPC server for the SecurityService
// service.
func NewSecurityServiceGatewayServer(svc SecurityServiceHandler, opts ...connect_gateway.HandlerOption) *SecurityServiceGatewayServer {
	return &SecurityServiceGatewayServer{
		listRoles:            connect_gateway.NewUnaryHandler(SecurityServiceListRolesProcedure, svc.ListRoles, opts...),
		createRole:           connect_gateway.NewUnaryHandler(SecurityServiceCreateRoleProcedure, svc.CreateRole, opts...),
		getRole:              connect_gateway.NewUnaryHandler(SecurityServiceGetRoleProcedure, svc.GetRole, opts...),
		updateRole:           connect_gateway.NewUnaryHandler(SecurityServiceUpdateRoleProcedure, svc.UpdateRole, opts...),
		deleteRole:           connect_gateway.NewUnaryHandler(SecurityServiceDeleteRoleProcedure, svc.DeleteRole, opts...),
		listRoleMembers:      connect_gateway.NewUnaryHandler(SecurityServiceListRoleMembersProcedure, svc.ListRoleMembers, opts...),
		updateRoleMembership: connect_gateway.NewUnaryHandler(SecurityServiceUpdateRoleMembershipProcedure, svc.UpdateRoleMembership, opts...),
		listRolesWithMembers: connect_gateway.NewUnaryHandler(SecurityServiceListRolesWithMembersProcedure, svc.ListRolesWithMembers, opts...),
	}
}

func (s *SecurityServiceGatewayServer) ListRoles(ctx context.Context, req *v1alpha1.ListRolesRequest) (*v1alpha1.ListRolesResponse, error) {
	return s.listRoles(ctx, req)
}

func (s *SecurityServiceGatewayServer) CreateRole(ctx context.Context, req *v1alpha1.CreateRoleRequest) (*v1alpha1.CreateRoleResponse, error) {
	return s.createRole(ctx, req)
}

func (s *SecurityServiceGatewayServer) GetRole(ctx context.Context, req *v1alpha1.GetRoleRequest) (*v1alpha1.GetRoleResponse, error) {
	return s.getRole(ctx, req)
}

func (s *SecurityServiceGatewayServer) UpdateRole(ctx context.Context, req *v1alpha1.UpdateRoleRequest) (*v1alpha1.UpdateRoleResponse, error) {
	return s.updateRole(ctx, req)
}

func (s *SecurityServiceGatewayServer) DeleteRole(ctx context.Context, req *v1alpha1.DeleteRoleRequest) (*v1alpha1.DeleteRoleResponse, error) {
	return s.deleteRole(ctx, req)
}

func (s *SecurityServiceGatewayServer) ListRoleMembers(ctx context.Context, req *v1alpha1.ListRoleMembersRequest) (*v1alpha1.ListRoleMembersResponse, error) {
	return s.listRoleMembers(ctx, req)
}

func (s *SecurityServiceGatewayServer) UpdateRoleMembership(ctx context.Context, req *v1alpha1.UpdateRoleMembershipRequest) (*v1alpha1.UpdateRoleMembershipResponse, error) {
	return s.updateRoleMembership(ctx, req)
}

func (s *SecurityServiceGatewayServer) ListRolesWithMembers(ctx context.Context, req *v1alpha1.ListRolesWithMembersRequest) (*v1alpha1.ListRolesWithMembersResponse, error) {
	return s.listRolesWithMembers(ctx, req)
}

// RegisterSecurityServiceHandlerGatewayServer registers the Connect handlers for the
// SecurityService "svc" to "mux".
func RegisterSecurityServiceHandlerGatewayServer(mux *runtime.ServeMux, svc SecurityServiceHandler, opts ...connect_gateway.HandlerOption) {
	if err := v1alpha1.RegisterSecurityServiceHandlerServer(context.TODO(), mux, NewSecurityServiceGatewayServer(svc, opts...)); err != nil {
		panic(fmt.Errorf("connect-gateway: %w", err))
	}
}
