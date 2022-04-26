// Code generated by MockGen. DO NOT EDIT.
// Source: wallet_grpc.pb.go

// Package mock___ is a generated GoMock package.
package mock___

import (
	context "context"
	__ "nn-blockchain-api/pkg/grpc_client/proto/wallet"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockWalletServiceClient is a mock of WalletServiceClient interface.
type MockWalletServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockWalletServiceClientMockRecorder
}

// MockWalletServiceClientMockRecorder is the mock recorder for MockWalletServiceClient.
type MockWalletServiceClientMockRecorder struct {
	mock *MockWalletServiceClient
}

// NewMockWalletServiceClient creates a new mock instance.
func NewMockWalletServiceClient(ctrl *gomock.Controller) *MockWalletServiceClient {
	mock := &MockWalletServiceClient{ctrl: ctrl}
	mock.recorder = &MockWalletServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWalletServiceClient) EXPECT() *MockWalletServiceClientMockRecorder {
	return m.recorder
}

// CreateMnemonic mocks base method.
func (m *MockWalletServiceClient) CreateMnemonic(ctx context.Context, in *__.CreateMnemonicData, opts ...grpc.CallOption) (*__.MnemonicInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMnemonic", varargs...)
	ret0, _ := ret[0].(*__.MnemonicInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMnemonic indicates an expected call of CreateMnemonic.
func (mr *MockWalletServiceClientMockRecorder) CreateMnemonic(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMnemonic", reflect.TypeOf((*MockWalletServiceClient)(nil).CreateMnemonic), varargs...)
}

// CreateWallet mocks base method.
func (m *MockWalletServiceClient) CreateWallet(ctx context.Context, in *__.CreateWalletData, opts ...grpc.CallOption) (*__.WalletInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateWallet", varargs...)
	ret0, _ := ret[0].(*__.WalletInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWallet indicates an expected call of CreateWallet.
func (mr *MockWalletServiceClientMockRecorder) CreateWallet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWallet", reflect.TypeOf((*MockWalletServiceClient)(nil).CreateWallet), varargs...)
}

// MockWalletServiceServer is a mock of WalletServiceServer interface.
type MockWalletServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockWalletServiceServerMockRecorder
}

// MockWalletServiceServerMockRecorder is the mock recorder for MockWalletServiceServer.
type MockWalletServiceServerMockRecorder struct {
	mock *MockWalletServiceServer
}

// NewMockWalletServiceServer creates a new mock instance.
func NewMockWalletServiceServer(ctrl *gomock.Controller) *MockWalletServiceServer {
	mock := &MockWalletServiceServer{ctrl: ctrl}
	mock.recorder = &MockWalletServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWalletServiceServer) EXPECT() *MockWalletServiceServerMockRecorder {
	return m.recorder
}

// CreateMnemonic mocks base method.
func (m *MockWalletServiceServer) CreateMnemonic(arg0 context.Context, arg1 *__.CreateMnemonicData) (*__.MnemonicInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMnemonic", arg0, arg1)
	ret0, _ := ret[0].(*__.MnemonicInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMnemonic indicates an expected call of CreateMnemonic.
func (mr *MockWalletServiceServerMockRecorder) CreateMnemonic(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMnemonic", reflect.TypeOf((*MockWalletServiceServer)(nil).CreateMnemonic), arg0, arg1)
}

// CreateWallet mocks base method.
func (m *MockWalletServiceServer) CreateWallet(arg0 context.Context, arg1 *__.CreateWalletData) (*__.WalletInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWallet", arg0, arg1)
	ret0, _ := ret[0].(*__.WalletInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWallet indicates an expected call of CreateWallet.
func (mr *MockWalletServiceServerMockRecorder) CreateWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWallet", reflect.TypeOf((*MockWalletServiceServer)(nil).CreateWallet), arg0, arg1)
}

// mustEmbedUnimplementedWalletServiceServer mocks base method.
func (m *MockWalletServiceServer) mustEmbedUnimplementedWalletServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedWalletServiceServer")
}

// mustEmbedUnimplementedWalletServiceServer indicates an expected call of mustEmbedUnimplementedWalletServiceServer.
func (mr *MockWalletServiceServerMockRecorder) mustEmbedUnimplementedWalletServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedWalletServiceServer", reflect.TypeOf((*MockWalletServiceServer)(nil).mustEmbedUnimplementedWalletServiceServer))
}

// MockUnsafeWalletServiceServer is a mock of UnsafeWalletServiceServer interface.
type MockUnsafeWalletServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeWalletServiceServerMockRecorder
}

// MockUnsafeWalletServiceServerMockRecorder is the mock recorder for MockUnsafeWalletServiceServer.
type MockUnsafeWalletServiceServerMockRecorder struct {
	mock *MockUnsafeWalletServiceServer
}

// NewMockUnsafeWalletServiceServer creates a new mock instance.
func NewMockUnsafeWalletServiceServer(ctrl *gomock.Controller) *MockUnsafeWalletServiceServer {
	mock := &MockUnsafeWalletServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeWalletServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeWalletServiceServer) EXPECT() *MockUnsafeWalletServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedWalletServiceServer mocks base method.
func (m *MockUnsafeWalletServiceServer) mustEmbedUnimplementedWalletServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedWalletServiceServer")
}

// mustEmbedUnimplementedWalletServiceServer indicates an expected call of mustEmbedUnimplementedWalletServiceServer.
func (mr *MockUnsafeWalletServiceServerMockRecorder) mustEmbedUnimplementedWalletServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedWalletServiceServer", reflect.TypeOf((*MockUnsafeWalletServiceServer)(nil).mustEmbedUnimplementedWalletServiceServer))
}
