// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gufranmirza/imdb-api/auth/jwt (interfaces: TokenAuth)

// Package jwtmock is a generated GoMock package.
package jwtmock

import (
	gomock "github.com/golang/mock/gomock"
	authmodel "github.com/gufranmirza/imdb-api/models/authmodel"
	http "net/http"
	reflect "reflect"
)

// MockTokenAuth is a mock of TokenAuth interface
type MockTokenAuth struct {
	ctrl     *gomock.Controller
	recorder *MockTokenAuthMockRecorder
}

// MockTokenAuthMockRecorder is the mock recorder for MockTokenAuth
type MockTokenAuthMockRecorder struct {
	mock *MockTokenAuth
}

// NewMockTokenAuth creates a new mock instance
func NewMockTokenAuth(ctrl *gomock.Controller) *MockTokenAuth {
	mock := &MockTokenAuth{ctrl: ctrl}
	mock.recorder = &MockTokenAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTokenAuth) EXPECT() *MockTokenAuthMockRecorder {
	return m.recorder
}

// CreateJWT mocks base method
func (m *MockTokenAuth) CreateJWT(arg0 authmodel.AppClaims) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJWT", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJWT indicates an expected call of CreateJWT
func (mr *MockTokenAuthMockRecorder) CreateJWT(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJWT", reflect.TypeOf((*MockTokenAuth)(nil).CreateJWT), arg0)
}

// CreateRefreshJWT mocks base method
func (m *MockTokenAuth) CreateRefreshJWT(arg0 authmodel.RefreshClaims) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRefreshJWT", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRefreshJWT indicates an expected call of CreateRefreshJWT
func (mr *MockTokenAuthMockRecorder) CreateRefreshJWT(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRefreshJWT", reflect.TypeOf((*MockTokenAuth)(nil).CreateRefreshJWT), arg0)
}

// GenTokenPair mocks base method
func (m *MockTokenAuth) GenTokenPair(arg0 authmodel.AppClaims, arg1 authmodel.RefreshClaims) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenTokenPair", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GenTokenPair indicates an expected call of GenTokenPair
func (mr *MockTokenAuthMockRecorder) GenTokenPair(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenTokenPair", reflect.TypeOf((*MockTokenAuth)(nil).GenTokenPair), arg0, arg1)
}

// Verifier mocks base method
func (m *MockTokenAuth) Verifier() func(http.Handler) http.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verifier")
	ret0, _ := ret[0].(func(http.Handler) http.Handler)
	return ret0
}

// Verifier indicates an expected call of Verifier
func (mr *MockTokenAuthMockRecorder) Verifier() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verifier", reflect.TypeOf((*MockTokenAuth)(nil).Verifier))
}
