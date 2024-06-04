// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/redpanda-data/console/backend/pkg/api (interfaces: AuthorizationHooks)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/authz_hooks.go -package=mocks . AuthorizationHooks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	rest "github.com/cloudhut/common/rest"
	gomock "go.uber.org/mock/gomock"
)

// MockAuthorizationHooks is a mock of AuthorizationHooks interface.
type MockAuthorizationHooks struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationHooksMockRecorder
}

// MockAuthorizationHooksMockRecorder is the mock recorder for MockAuthorizationHooks.
type MockAuthorizationHooksMockRecorder struct {
	mock *MockAuthorizationHooks
}

// NewMockAuthorizationHooks creates a new mock instance.
func NewMockAuthorizationHooks(ctrl *gomock.Controller) *MockAuthorizationHooks {
	mock := &MockAuthorizationHooks{ctrl: ctrl}
	mock.recorder = &MockAuthorizationHooksMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizationHooks) EXPECT() *MockAuthorizationHooksMockRecorder {
	return m.recorder
}

// AllowedConnectClusterActions mocks base method.
func (m *MockAuthorizationHooks) AllowedConnectClusterActions(arg0 context.Context, arg1 string) ([]string, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllowedConnectClusterActions", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// AllowedConnectClusterActions indicates an expected call of AllowedConnectClusterActions.
func (mr *MockAuthorizationHooksMockRecorder) AllowedConnectClusterActions(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllowedConnectClusterActions", reflect.TypeOf((*MockAuthorizationHooks)(nil).AllowedConnectClusterActions), arg0, arg1)
}

// AllowedConsumerGroupActions mocks base method.
func (m *MockAuthorizationHooks) AllowedConsumerGroupActions(arg0 context.Context, arg1 string) ([]string, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllowedConsumerGroupActions", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// AllowedConsumerGroupActions indicates an expected call of AllowedConsumerGroupActions.
func (mr *MockAuthorizationHooksMockRecorder) AllowedConsumerGroupActions(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllowedConsumerGroupActions", reflect.TypeOf((*MockAuthorizationHooks)(nil).AllowedConsumerGroupActions), arg0, arg1)
}

// AllowedTopicActions mocks base method.
func (m *MockAuthorizationHooks) AllowedTopicActions(arg0 context.Context, arg1 string) ([]string, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllowedTopicActions", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// AllowedTopicActions indicates an expected call of AllowedTopicActions.
func (mr *MockAuthorizationHooksMockRecorder) AllowedTopicActions(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllowedTopicActions", reflect.TypeOf((*MockAuthorizationHooks)(nil).AllowedTopicActions), arg0, arg1)
}

// CanCreateACL mocks base method.
func (m *MockAuthorizationHooks) CanCreateACL(arg0 context.Context) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanCreateACL", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanCreateACL indicates an expected call of CanCreateACL.
func (mr *MockAuthorizationHooksMockRecorder) CanCreateACL(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanCreateACL", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanCreateACL), arg0)
}

// CanCreateKafkaUsers mocks base method.
func (m *MockAuthorizationHooks) CanCreateKafkaUsers(arg0 context.Context) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanCreateKafkaUsers", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanCreateKafkaUsers indicates an expected call of CanCreateKafkaUsers.
func (mr *MockAuthorizationHooksMockRecorder) CanCreateKafkaUsers(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanCreateKafkaUsers", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanCreateKafkaUsers), arg0)
}

// CanCreateRedpandaRoles mocks base method.
func (m *MockAuthorizationHooks) CanCreateRedpandaRoles(arg0 context.Context) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanCreateRedpandaRoles", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanCreateRedpandaRoles indicates an expected call of CanCreateRedpandaRoles.
func (mr *MockAuthorizationHooksMockRecorder) CanCreateRedpandaRoles(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanCreateRedpandaRoles", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanCreateRedpandaRoles), arg0)
}

// CanCreateSchemas mocks base method.
func (m *MockAuthorizationHooks) CanCreateSchemas(arg0 context.Context) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanCreateSchemas", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanCreateSchemas indicates an expected call of CanCreateSchemas.
func (mr *MockAuthorizationHooksMockRecorder) CanCreateSchemas(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanCreateSchemas", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanCreateSchemas), arg0)
}

// CanCreateTopic mocks base method.
func (m *MockAuthorizationHooks) CanCreateTopic(arg0 context.Context, arg1 string) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanCreateTopic", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanCreateTopic indicates an expected call of CanCreateTopic.
func (mr *MockAuthorizationHooksMockRecorder) CanCreateTopic(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanCreateTopic", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanCreateTopic), arg0, arg1)
}

// CanDeleteACL mocks base method.
func (m *MockAuthorizationHooks) CanDeleteACL(arg0 context.Context) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanDeleteACL", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanDeleteACL indicates an expected call of CanDeleteACL.
func (mr *MockAuthorizationHooksMockRecorder) CanDeleteACL(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanDeleteACL", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanDeleteACL), arg0)
}

// CanDeleteConnectCluster mocks base method.
func (m *MockAuthorizationHooks) CanDeleteConnectCluster(arg0 context.Context, arg1 string) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanDeleteConnectCluster", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanDeleteConnectCluster indicates an expected call of CanDeleteConnectCluster.
func (mr *MockAuthorizationHooksMockRecorder) CanDeleteConnectCluster(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanDeleteConnectCluster", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanDeleteConnectCluster), arg0, arg1)
}

// CanDeleteConsumerGroup mocks base method.
func (m *MockAuthorizationHooks) CanDeleteConsumerGroup(arg0 context.Context, arg1 string) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanDeleteConsumerGroup", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanDeleteConsumerGroup indicates an expected call of CanDeleteConsumerGroup.
func (mr *MockAuthorizationHooksMockRecorder) CanDeleteConsumerGroup(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanDeleteConsumerGroup", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanDeleteConsumerGroup), arg0, arg1)
}

// CanDeleteKafkaUsers mocks base method.
func (m *MockAuthorizationHooks) CanDeleteKafkaUsers(arg0 context.Context) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanDeleteKafkaUsers", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanDeleteKafkaUsers indicates an expected call of CanDeleteKafkaUsers.
func (mr *MockAuthorizationHooksMockRecorder) CanDeleteKafkaUsers(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanDeleteKafkaUsers", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanDeleteKafkaUsers), arg0)
}

// CanDeleteRedpandaRoles mocks base method.
func (m *MockAuthorizationHooks) CanDeleteRedpandaRoles(arg0 context.Context) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanDeleteRedpandaRoles", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanDeleteRedpandaRoles indicates an expected call of CanDeleteRedpandaRoles.
func (mr *MockAuthorizationHooksMockRecorder) CanDeleteRedpandaRoles(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanDeleteRedpandaRoles", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanDeleteRedpandaRoles), arg0)
}

// CanDeleteSchemas mocks base method.
func (m *MockAuthorizationHooks) CanDeleteSchemas(arg0 context.Context) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanDeleteSchemas", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanDeleteSchemas indicates an expected call of CanDeleteSchemas.
func (mr *MockAuthorizationHooksMockRecorder) CanDeleteSchemas(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanDeleteSchemas", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanDeleteSchemas), arg0)
}

// CanDeleteTopic mocks base method.
func (m *MockAuthorizationHooks) CanDeleteTopic(arg0 context.Context, arg1 string) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanDeleteTopic", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanDeleteTopic indicates an expected call of CanDeleteTopic.
func (mr *MockAuthorizationHooksMockRecorder) CanDeleteTopic(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanDeleteTopic", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanDeleteTopic), arg0, arg1)
}

// CanDeleteTopicRecords mocks base method.
func (m *MockAuthorizationHooks) CanDeleteTopicRecords(arg0 context.Context, arg1 string) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanDeleteTopicRecords", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanDeleteTopicRecords indicates an expected call of CanDeleteTopicRecords.
func (mr *MockAuthorizationHooksMockRecorder) CanDeleteTopicRecords(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanDeleteTopicRecords", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanDeleteTopicRecords), arg0, arg1)
}

// CanEditConnectCluster mocks base method.
func (m *MockAuthorizationHooks) CanEditConnectCluster(arg0 context.Context, arg1 string) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanEditConnectCluster", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanEditConnectCluster indicates an expected call of CanEditConnectCluster.
func (mr *MockAuthorizationHooksMockRecorder) CanEditConnectCluster(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanEditConnectCluster", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanEditConnectCluster), arg0, arg1)
}

// CanEditConsumerGroup mocks base method.
func (m *MockAuthorizationHooks) CanEditConsumerGroup(arg0 context.Context, arg1 string) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanEditConsumerGroup", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanEditConsumerGroup indicates an expected call of CanEditConsumerGroup.
func (mr *MockAuthorizationHooksMockRecorder) CanEditConsumerGroup(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanEditConsumerGroup", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanEditConsumerGroup), arg0, arg1)
}

// CanEditTopicConfig mocks base method.
func (m *MockAuthorizationHooks) CanEditTopicConfig(arg0 context.Context, arg1 string) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanEditTopicConfig", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanEditTopicConfig indicates an expected call of CanEditTopicConfig.
func (mr *MockAuthorizationHooksMockRecorder) CanEditTopicConfig(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanEditTopicConfig", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanEditTopicConfig), arg0, arg1)
}

// CanListACLs mocks base method.
func (m *MockAuthorizationHooks) CanListACLs(arg0 context.Context) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanListACLs", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanListACLs indicates an expected call of CanListACLs.
func (mr *MockAuthorizationHooksMockRecorder) CanListACLs(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanListACLs", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanListACLs), arg0)
}

// CanListKafkaUsers mocks base method.
func (m *MockAuthorizationHooks) CanListKafkaUsers(arg0 context.Context) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanListKafkaUsers", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanListKafkaUsers indicates an expected call of CanListKafkaUsers.
func (mr *MockAuthorizationHooksMockRecorder) CanListKafkaUsers(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanListKafkaUsers", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanListKafkaUsers), arg0)
}

// CanListQuotas mocks base method.
func (m *MockAuthorizationHooks) CanListQuotas(arg0 context.Context) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanListQuotas", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanListQuotas indicates an expected call of CanListQuotas.
func (mr *MockAuthorizationHooksMockRecorder) CanListQuotas(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanListQuotas", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanListQuotas), arg0)
}

// CanListRedpandaRoles mocks base method.
func (m *MockAuthorizationHooks) CanListRedpandaRoles(arg0 context.Context) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanListRedpandaRoles", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanListRedpandaRoles indicates an expected call of CanListRedpandaRoles.
func (mr *MockAuthorizationHooksMockRecorder) CanListRedpandaRoles(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanListRedpandaRoles", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanListRedpandaRoles), arg0)
}

// CanManageSchemaRegistry mocks base method.
func (m *MockAuthorizationHooks) CanManageSchemaRegistry(arg0 context.Context) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanManageSchemaRegistry", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanManageSchemaRegistry indicates an expected call of CanManageSchemaRegistry.
func (mr *MockAuthorizationHooksMockRecorder) CanManageSchemaRegistry(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanManageSchemaRegistry", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanManageSchemaRegistry), arg0)
}

// CanPatchConfigs mocks base method.
func (m *MockAuthorizationHooks) CanPatchConfigs(arg0 context.Context) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanPatchConfigs", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanPatchConfigs indicates an expected call of CanPatchConfigs.
func (mr *MockAuthorizationHooksMockRecorder) CanPatchConfigs(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanPatchConfigs", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanPatchConfigs), arg0)
}

// CanPatchPartitionReassignments mocks base method.
func (m *MockAuthorizationHooks) CanPatchPartitionReassignments(arg0 context.Context) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanPatchPartitionReassignments", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanPatchPartitionReassignments indicates an expected call of CanPatchPartitionReassignments.
func (mr *MockAuthorizationHooksMockRecorder) CanPatchPartitionReassignments(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanPatchPartitionReassignments", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanPatchPartitionReassignments), arg0)
}

// CanSeeConsumerGroup mocks base method.
func (m *MockAuthorizationHooks) CanSeeConsumerGroup(arg0 context.Context, arg1 string) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanSeeConsumerGroup", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanSeeConsumerGroup indicates an expected call of CanSeeConsumerGroup.
func (mr *MockAuthorizationHooksMockRecorder) CanSeeConsumerGroup(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanSeeConsumerGroup", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanSeeConsumerGroup), arg0, arg1)
}

// CanSeeTopic mocks base method.
func (m *MockAuthorizationHooks) CanSeeTopic(arg0 context.Context, arg1 string) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanSeeTopic", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanSeeTopic indicates an expected call of CanSeeTopic.
func (mr *MockAuthorizationHooksMockRecorder) CanSeeTopic(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanSeeTopic", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanSeeTopic), arg0, arg1)
}

// CanViewConnectCluster mocks base method.
func (m *MockAuthorizationHooks) CanViewConnectCluster(arg0 context.Context, arg1 string) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanViewConnectCluster", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanViewConnectCluster indicates an expected call of CanViewConnectCluster.
func (mr *MockAuthorizationHooksMockRecorder) CanViewConnectCluster(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanViewConnectCluster", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanViewConnectCluster), arg0, arg1)
}

// CanViewSchemas mocks base method.
func (m *MockAuthorizationHooks) CanViewSchemas(arg0 context.Context) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanViewSchemas", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanViewSchemas indicates an expected call of CanViewSchemas.
func (mr *MockAuthorizationHooksMockRecorder) CanViewSchemas(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanViewSchemas", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanViewSchemas), arg0)
}

// CanViewTopicConfig mocks base method.
func (m *MockAuthorizationHooks) CanViewTopicConfig(arg0 context.Context, arg1 string) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanViewTopicConfig", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanViewTopicConfig indicates an expected call of CanViewTopicConfig.
func (mr *MockAuthorizationHooksMockRecorder) CanViewTopicConfig(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanViewTopicConfig", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanViewTopicConfig), arg0, arg1)
}

// CanViewTopicConsumers mocks base method.
func (m *MockAuthorizationHooks) CanViewTopicConsumers(arg0 context.Context, arg1 string) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanViewTopicConsumers", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanViewTopicConsumers indicates an expected call of CanViewTopicConsumers.
func (mr *MockAuthorizationHooksMockRecorder) CanViewTopicConsumers(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanViewTopicConsumers", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanViewTopicConsumers), arg0, arg1)
}

// CanViewTopicPartitions mocks base method.
func (m *MockAuthorizationHooks) CanViewTopicPartitions(arg0 context.Context, arg1 string) (bool, *rest.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanViewTopicPartitions", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*rest.Error)
	return ret0, ret1
}

// CanViewTopicPartitions indicates an expected call of CanViewTopicPartitions.
func (mr *MockAuthorizationHooksMockRecorder) CanViewTopicPartitions(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanViewTopicPartitions", reflect.TypeOf((*MockAuthorizationHooks)(nil).CanViewTopicPartitions), arg0, arg1)
}

// IsProtectedKafkaUser mocks base method.
func (m *MockAuthorizationHooks) IsProtectedKafkaUser(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsProtectedKafkaUser", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsProtectedKafkaUser indicates an expected call of IsProtectedKafkaUser.
func (mr *MockAuthorizationHooksMockRecorder) IsProtectedKafkaUser(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsProtectedKafkaUser", reflect.TypeOf((*MockAuthorizationHooks)(nil).IsProtectedKafkaUser), arg0)
}
