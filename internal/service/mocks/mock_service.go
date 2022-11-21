// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	core "coins-app/internal/core"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthorization) CreateUser(user core.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), user)
}

// GenerateToken mocks base method.
func (m *MockAuthorization) GenerateToken(username, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", username, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthorizationMockRecorder) GenerateToken(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), username, password)
}

// ParseToken mocks base method.
func (m *MockAuthorization) ParseToken(token string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthorizationMockRecorder) ParseToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthorization)(nil).ParseToken), token)
}

// MockAccount is a mock of Account interface.
type MockAccount struct {
	ctrl     *gomock.Controller
	recorder *MockAccountMockRecorder
}

// MockAccountMockRecorder is the mock recorder for MockAccount.
type MockAccountMockRecorder struct {
	mock *MockAccount
}

// NewMockAccount creates a new mock instance.
func NewMockAccount(ctrl *gomock.Controller) *MockAccount {
	mock := &MockAccount{ctrl: ctrl}
	mock.recorder = &MockAccountMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccount) EXPECT() *MockAccountMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockAccount) CreateAccount(account core.Account) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", account)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockAccountMockRecorder) CreateAccount(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockAccount)(nil).CreateAccount), account)
}

// GetAccountById mocks base method.
func (m *MockAccount) GetAccountById(accountId int) (core.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountById", accountId)
	ret0, _ := ret[0].(core.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountById indicates an expected call of GetAccountById.
func (mr *MockAccountMockRecorder) GetAccountById(accountId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountById", reflect.TypeOf((*MockAccount)(nil).GetAccountById), accountId)
}

// GetAccounts mocks base method.
func (m *MockAccount) GetAccounts(userId int) ([]core.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccounts", userId)
	ret0, _ := ret[0].([]core.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccounts indicates an expected call of GetAccounts.
func (mr *MockAccountMockRecorder) GetAccounts(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*MockAccount)(nil).GetAccounts), userId)
}

// UpdateAccount mocks base method.
func (m *MockAccount) UpdateAccount(account core.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", account)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockAccountMockRecorder) UpdateAccount(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockAccount)(nil).UpdateAccount), account)
}

// MockTransfer is a mock of Transfer interface.
type MockTransfer struct {
	ctrl     *gomock.Controller
	recorder *MockTransferMockRecorder
}

// MockTransferMockRecorder is the mock recorder for MockTransfer.
type MockTransferMockRecorder struct {
	mock *MockTransfer
}

// NewMockTransfer creates a new mock instance.
func NewMockTransfer(ctrl *gomock.Controller) *MockTransfer {
	mock := &MockTransfer{ctrl: ctrl}
	mock.recorder = &MockTransferMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransfer) EXPECT() *MockTransferMockRecorder {
	return m.recorder
}

// CreateTransfer mocks base method.
func (m *MockTransfer) CreateTransfer(transfer core.Transfer) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransfer", transfer)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransfer indicates an expected call of CreateTransfer.
func (mr *MockTransferMockRecorder) CreateTransfer(transfer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransfer", reflect.TypeOf((*MockTransfer)(nil).CreateTransfer), transfer)
}

// GetTransferById mocks base method.
func (m *MockTransfer) GetTransferById(transferId int) (core.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransferById", transferId)
	ret0, _ := ret[0].(core.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransferById indicates an expected call of GetTransferById.
func (mr *MockTransferMockRecorder) GetTransferById(transferId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransferById", reflect.TypeOf((*MockTransfer)(nil).GetTransferById), transferId)
}

// GetTransfers mocks base method.
func (m *MockTransfer) GetTransfers(userId int) ([]core.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransfers", userId)
	ret0, _ := ret[0].([]core.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransfers indicates an expected call of GetTransfers.
func (mr *MockTransferMockRecorder) GetTransfers(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransfers", reflect.TypeOf((*MockTransfer)(nil).GetTransfers), userId)
}

// MockCoin is a mock of Coin interface.
type MockCoin struct {
	ctrl     *gomock.Controller
	recorder *MockCoinMockRecorder
}

// MockCoinMockRecorder is the mock recorder for MockCoin.
type MockCoinMockRecorder struct {
	mock *MockCoin
}

// NewMockCoin creates a new mock instance.
func NewMockCoin(ctrl *gomock.Controller) *MockCoin {
	mock := &MockCoin{ctrl: ctrl}
	mock.recorder = &MockCoinMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCoin) EXPECT() *MockCoinMockRecorder {
	return m.recorder
}

// GetCoinPrices mocks base method.
func (m *MockCoin) GetCoinPrices(symbol string) ([]core.SymbolPrice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCoinPrices", symbol)
	ret0, _ := ret[0].([]core.SymbolPrice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCoinPrices indicates an expected call of GetCoinPrices.
func (mr *MockCoinMockRecorder) GetCoinPrices(symbol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCoinPrices", reflect.TypeOf((*MockCoin)(nil).GetCoinPrices), symbol)
}
