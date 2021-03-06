// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_bitcoin_rpc is a generated GoMock package.
package mock_bitcoin_rpc

import (
	context "context"
	bitcoin_rpc "nn-blockchain-api/pkg/rpc/bitcoin"
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
func (m *MockService) CreateTransaction(ctx context.Context, utxos bitcoin_rpc.UTXO, fromAddress, toAddress string, amount int64, network string) (*string, *float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", ctx, utxos, fromAddress, toAddress, amount, network)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(*float64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockServiceMockRecorder) CreateTransaction(ctx, utxos, fromAddress, toAddress, amount, network interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockService)(nil).CreateTransaction), ctx, utxos, fromAddress, toAddress, amount, network)
}

// CreateWallet mocks base method.
func (m *MockService) CreateWallet(ctx context.Context, network string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWallet", ctx, network)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWallet indicates an expected call of CreateWallet.
func (mr *MockServiceMockRecorder) CreateWallet(ctx, network interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWallet", reflect.TypeOf((*MockService)(nil).CreateWallet), ctx, network)
}

// DecodeTransaction mocks base method.
func (m *MockService) DecodeTransaction(ctx context.Context, tx, network string) (*bitcoin_rpc.DecodedTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecodeTransaction", ctx, tx, network)
	ret0, _ := ret[0].(*bitcoin_rpc.DecodedTx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecodeTransaction indicates an expected call of DecodeTransaction.
func (mr *MockServiceMockRecorder) DecodeTransaction(ctx, tx, network interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeTransaction", reflect.TypeOf((*MockService)(nil).DecodeTransaction), ctx, tx, network)
}

// FundForTransaction mocks base method.
func (m *MockService) FundForTransaction(ctx context.Context, createdTx, changeAddress, network string) (string, *float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FundForTransaction", ctx, createdTx, changeAddress, network)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*float64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FundForTransaction indicates an expected call of FundForTransaction.
func (mr *MockServiceMockRecorder) FundForTransaction(ctx, createdTx, changeAddress, network interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FundForTransaction", reflect.TypeOf((*MockService)(nil).FundForTransaction), ctx, createdTx, changeAddress, network)
}

// GetCurrentFee mocks base method.
func (m *MockService) GetCurrentFee(ctx context.Context, network string) (*float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentFee", ctx, network)
	ret0, _ := ret[0].(*float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentFee indicates an expected call of GetCurrentFee.
func (mr *MockServiceMockRecorder) GetCurrentFee(ctx, network interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentFee", reflect.TypeOf((*MockService)(nil).GetCurrentFee), ctx, network)
}

// ImportAddress mocks base method.
func (m *MockService) ImportAddress(ctx context.Context, address, walletId, network string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportAddress", ctx, address, walletId, network)
	ret0, _ := ret[0].(error)
	return ret0
}

// ImportAddress indicates an expected call of ImportAddress.
func (mr *MockServiceMockRecorder) ImportAddress(ctx, address, walletId, network interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportAddress", reflect.TypeOf((*MockService)(nil).ImportAddress), ctx, address, walletId, network)
}

// ListUnspent mocks base method.
func (m *MockService) ListUnspent(ctx context.Context, address, walletId, network string) ([]*bitcoin_rpc.Unspent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUnspent", ctx, address, walletId, network)
	ret0, _ := ret[0].([]*bitcoin_rpc.Unspent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUnspent indicates an expected call of ListUnspent.
func (mr *MockServiceMockRecorder) ListUnspent(ctx, address, walletId, network interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUnspent", reflect.TypeOf((*MockService)(nil).ListUnspent), ctx, address, walletId, network)
}

// LoadWallet mocks base method.
func (m *MockService) LoadWallet(ctx context.Context, walletId, network string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadWallet", ctx, walletId, network)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadWallet indicates an expected call of LoadWallet.
func (mr *MockServiceMockRecorder) LoadWallet(ctx, walletId, network interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadWallet", reflect.TypeOf((*MockService)(nil).LoadWallet), ctx, walletId, network)
}

// RescanWallet mocks base method.
func (m *MockService) RescanWallet(ctx context.Context, walletId, network string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RescanWallet", ctx, walletId, network)
	ret0, _ := ret[0].(error)
	return ret0
}

// RescanWallet indicates an expected call of RescanWallet.
func (mr *MockServiceMockRecorder) RescanWallet(ctx, walletId, network interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RescanWallet", reflect.TypeOf((*MockService)(nil).RescanWallet), ctx, walletId, network)
}

// SendTransaction mocks base method.
func (m *MockService) SendTransaction(ctx context.Context, signedTx, network string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransaction", ctx, signedTx, network)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTransaction indicates an expected call of SendTransaction.
func (mr *MockServiceMockRecorder) SendTransaction(ctx, signedTx, network interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransaction", reflect.TypeOf((*MockService)(nil).SendTransaction), ctx, signedTx, network)
}

// SignTransaction mocks base method.
func (m *MockService) SignTransaction(ctx context.Context, tx, privateKey string, utxos bitcoin_rpc.UTXO, network string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignTransaction", ctx, tx, privateKey, utxos, network)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignTransaction indicates an expected call of SignTransaction.
func (mr *MockServiceMockRecorder) SignTransaction(ctx, tx, privateKey, utxos, network interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignTransaction", reflect.TypeOf((*MockService)(nil).SignTransaction), ctx, tx, privateKey, utxos, network)
}

// Status mocks base method.
func (m *MockService) Status(ctx context.Context, network string) (*bitcoin_rpc.StatusNode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", ctx, network)
	ret0, _ := ret[0].(*bitcoin_rpc.StatusNode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockServiceMockRecorder) Status(ctx, network interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockService)(nil).Status), ctx, network)
}

// WalletInfo mocks base method.
func (m *MockService) WalletInfo(ctx context.Context, walletId, network string) (*bitcoin_rpc.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletInfo", ctx, walletId, network)
	ret0, _ := ret[0].(*bitcoin_rpc.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletInfo indicates an expected call of WalletInfo.
func (mr *MockServiceMockRecorder) WalletInfo(ctx, walletId, network interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletInfo", reflect.TypeOf((*MockService)(nil).WalletInfo), ctx, walletId, network)
}
