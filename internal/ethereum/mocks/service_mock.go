// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_ethereum is a generated GoMock package.
package mock_ethereum

import (
	context "context"
	ethereum "nn-blockchain-api/internal/ethereum"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateTransaction mocks base method.
func (m *MockService) CreateTransaction(ctx context.Context, dto *ethereum.CreateRawTransactionDTO) (*ethereum.CreatedRawTransactionDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", ctx, dto)
	ret0, _ := ret[0].(*ethereum.CreatedRawTransactionDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockServiceMockRecorder) CreateTransaction(ctx, dto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockService)(nil).CreateTransaction), ctx, dto)
}

// SendTransaction mocks base method.
func (m *MockService) SendTransaction(ctx context.Context, dto *ethereum.SendRawTransactionDTO) (*ethereum.SentRawTransactionDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransaction", ctx, dto)
	ret0, _ := ret[0].(*ethereum.SentRawTransactionDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTransaction indicates an expected call of SendTransaction.
func (mr *MockServiceMockRecorder) SendTransaction(ctx, dto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransaction", reflect.TypeOf((*MockService)(nil).SendTransaction), ctx, dto)
}

// SignTransaction mocks base method.
func (m *MockService) SignTransaction(ctx context.Context, dto *ethereum.SignRawTransactionDTO) (*ethereum.SignedRawTransactionDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignTransaction", ctx, dto)
	ret0, _ := ret[0].(*ethereum.SignedRawTransactionDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignTransaction indicates an expected call of SignTransaction.
func (mr *MockServiceMockRecorder) SignTransaction(ctx, dto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignTransaction", reflect.TypeOf((*MockService)(nil).SignTransaction), ctx, dto)
}

// StatusNode mocks base method.
func (m *MockService) StatusNode(ctx context.Context, dto *ethereum.StatusNodeDTO) (*ethereum.NodeInfoDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatusNode", ctx, dto)
	ret0, _ := ret[0].(*ethereum.NodeInfoDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatusNode indicates an expected call of StatusNode.
func (mr *MockServiceMockRecorder) StatusNode(ctx, dto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusNode", reflect.TypeOf((*MockService)(nil).StatusNode), ctx, dto)
}
