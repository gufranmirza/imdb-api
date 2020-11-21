// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gufranmirza/imdb-api/dal/tokendal (interfaces: TokenDal)

// Package tokendalmock is a generated GoMock package.
package tokendalmock

import (
	gomock "github.com/golang/mock/gomock"
	dbmodels "github.com/gufranmirza/imdb-api/database/dbmodels"
	reflect "reflect"
)

// MockTokenDal is a mock of TokenDal interface
type MockTokenDal struct {
	ctrl     *gomock.Controller
	recorder *MockTokenDalMockRecorder
}

// MockTokenDalMockRecorder is the mock recorder for MockTokenDal
type MockTokenDalMockRecorder struct {
	mock *MockTokenDal
}

// NewMockTokenDal creates a new mock instance
func NewMockTokenDal(ctrl *gomock.Controller) *MockTokenDal {
	mock := &MockTokenDal{ctrl: ctrl}
	mock.recorder = &MockTokenDalMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTokenDal) EXPECT() *MockTokenDalMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockTokenDal) Create(arg0 string, arg1 *dbmodels.Token) (*dbmodels.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*dbmodels.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockTokenDalMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTokenDal)(nil).Create), arg0, arg1)
}

// DeleteByAccessToken mocks base method
func (m *MockTokenDal) DeleteByAccessToken(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByAccessToken", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByAccessToken indicates an expected call of DeleteByAccessToken
func (mr *MockTokenDalMockRecorder) DeleteByAccessToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByAccessToken", reflect.TypeOf((*MockTokenDal)(nil).DeleteByAccessToken), arg0)
}

// GetByUUID mocks base method
func (m *MockTokenDal) GetByUUID(arg0 string) (*dbmodels.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUUID", arg0)
	ret0, _ := ret[0].(*dbmodels.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUUID indicates an expected call of GetByUUID
func (mr *MockTokenDalMockRecorder) GetByUUID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUUID", reflect.TypeOf((*MockTokenDal)(nil).GetByUUID), arg0)
}

// Update mocks base method
func (m *MockTokenDal) Update(arg0 *dbmodels.Token) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockTokenDalMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTokenDal)(nil).Update), arg0)
}