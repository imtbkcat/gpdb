// Code generated by MockGen. DO NOT EDIT.
// Source: idl/command.pb.go

// Package mock_idl is a generated GoMock package.
package mock_idl

import (
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	. "gp_upgrade/idl"
	reflect "reflect"
)

// MockCommandListenerClient is a mock of CommandListenerClient interface
type MockCommandListenerClient struct {
	ctrl     *gomock.Controller
	recorder *MockCommandListenerClientMockRecorder
}

// MockCommandListenerClientMockRecorder is the mock recorder for MockCommandListenerClient
type MockCommandListenerClientMockRecorder struct {
	mock *MockCommandListenerClient
}

// NewMockCommandListenerClient creates a new mock instance
func NewMockCommandListenerClient(ctrl *gomock.Controller) *MockCommandListenerClient {
	mock := &MockCommandListenerClient{ctrl: ctrl}
	mock.recorder = &MockCommandListenerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommandListenerClient) EXPECT() *MockCommandListenerClientMockRecorder {
	return m.recorder
}

// CheckUpgradeStatus mocks base method
func (m *MockCommandListenerClient) CheckUpgradeStatus(ctx context.Context, in *CheckUpgradeStatusRequest, opts ...grpc.CallOption) (*CheckUpgradeStatusReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckUpgradeStatus", varargs...)
	ret0, _ := ret[0].(*CheckUpgradeStatusReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUpgradeStatus indicates an expected call of CheckUpgradeStatus
func (mr *MockCommandListenerClientMockRecorder) CheckUpgradeStatus(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUpgradeStatus", reflect.TypeOf((*MockCommandListenerClient)(nil).CheckUpgradeStatus), varargs...)
}

// CheckDiskUsage mocks base method
func (m *MockCommandListenerClient) CheckDiskUsage(ctx context.Context, in *CheckDiskUsageRequest, opts ...grpc.CallOption) (*CheckDiskUsageReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckDiskUsage", varargs...)
	ret0, _ := ret[0].(*CheckDiskUsageReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckDiskUsage indicates an expected call of CheckDiskUsage
func (mr *MockCommandListenerClientMockRecorder) CheckDiskUsage(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDiskUsage", reflect.TypeOf((*MockCommandListenerClient)(nil).CheckDiskUsage), varargs...)
}

// MockCommandListenerServer is a mock of CommandListenerServer interface
type MockCommandListenerServer struct {
	ctrl     *gomock.Controller
	recorder *MockCommandListenerServerMockRecorder
}

// MockCommandListenerServerMockRecorder is the mock recorder for MockCommandListenerServer
type MockCommandListenerServerMockRecorder struct {
	mock *MockCommandListenerServer
}

// NewMockCommandListenerServer creates a new mock instance
func NewMockCommandListenerServer(ctrl *gomock.Controller) *MockCommandListenerServer {
	mock := &MockCommandListenerServer{ctrl: ctrl}
	mock.recorder = &MockCommandListenerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommandListenerServer) EXPECT() *MockCommandListenerServerMockRecorder {
	return m.recorder
}

// CheckUpgradeStatus mocks base method
func (m *MockCommandListenerServer) CheckUpgradeStatus(arg0 context.Context, arg1 *CheckUpgradeStatusRequest) (*CheckUpgradeStatusReply, error) {
	ret := m.ctrl.Call(m, "CheckUpgradeStatus", arg0, arg1)
	ret0, _ := ret[0].(*CheckUpgradeStatusReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUpgradeStatus indicates an expected call of CheckUpgradeStatus
func (mr *MockCommandListenerServerMockRecorder) CheckUpgradeStatus(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUpgradeStatus", reflect.TypeOf((*MockCommandListenerServer)(nil).CheckUpgradeStatus), arg0, arg1)
}

// CheckDiskUsage mocks base method
func (m *MockCommandListenerServer) CheckDiskUsage(arg0 context.Context, arg1 *CheckDiskUsageRequest) (*CheckDiskUsageReply, error) {
	ret := m.ctrl.Call(m, "CheckDiskUsage", arg0, arg1)
	ret0, _ := ret[0].(*CheckDiskUsageReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckDiskUsage indicates an expected call of CheckDiskUsage
func (mr *MockCommandListenerServerMockRecorder) CheckDiskUsage(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDiskUsage", reflect.TypeOf((*MockCommandListenerServer)(nil).CheckDiskUsage), arg0, arg1)
}