// Code generated by MockGen. DO NOT EDIT.
// Source: health.go

// Package mock_rpc_bitcoin is a generated GoMock package.
package mock_rpc_bitcoin

import (
	context "context"
	rpc_bitcoin "nn-blockchain-api/pkg/rpc/bitcoin"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHealthService is a mock of HealthService interface.
type MockHealthService struct {
	ctrl     *gomock.Controller
	recorder *MockHealthServiceMockRecorder
}

// MockHealthServiceMockRecorder is the mock recorder for MockHealthService.
type MockHealthServiceMockRecorder struct {
	mock *MockHealthService
}

// NewMockHealthService creates a new mock instance.
func NewMockHealthService(ctrl *gomock.Controller) *MockHealthService {
	mock := &MockHealthService{ctrl: ctrl}
	mock.recorder = &MockHealthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHealthService) EXPECT() *MockHealthServiceMockRecorder {
	return m.recorder
}

// Status mocks base method.
func (m *MockHealthService) Status(ctx context.Context, network string) (*rpc_bitcoin.StatusNode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", ctx, network)
	ret0, _ := ret[0].(*rpc_bitcoin.StatusNode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockHealthServiceMockRecorder) Status(ctx, network interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockHealthService)(nil).Status), ctx, network)
}
