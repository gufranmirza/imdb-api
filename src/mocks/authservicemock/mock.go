// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gufranmirza/imdb-api/web/services/v1/authservice (interfaces: AuthService)

// Package authservicemock is a generated GoMock package.
package authservicemock

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockAuthService is a mock of AuthService interface
type MockAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService
type MockAuthServiceMockRecorder struct {
	mock *MockAuthService
}

// NewMockAuthService creates a new mock instance
func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
	mock := &MockAuthService{ctrl: ctrl}
	mock.recorder = &MockAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
	return m.recorder
}

// Login mocks base method
func (m *MockAuthService) Login(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Login", arg0, arg1)
}

// Login indicates an expected call of Login
func (mr *MockAuthServiceMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthService)(nil).Login), arg0, arg1)
}

// Logout mocks base method
func (m *MockAuthService) Logout(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Logout", arg0, arg1)
}

// Logout indicates an expected call of Logout
func (mr *MockAuthServiceMockRecorder) Logout(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAuthService)(nil).Logout), arg0, arg1)
}

// SignUp mocks base method
func (m *MockAuthService) SignUp(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SignUp", arg0, arg1)
}

// SignUp indicates an expected call of SignUp
func (mr *MockAuthServiceMockRecorder) SignUp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockAuthService)(nil).SignUp), arg0, arg1)
}

// Validate mocks base method
func (m *MockAuthService) Validate(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Validate", arg0, arg1)
}

// Validate indicates an expected call of Validate
func (mr *MockAuthServiceMockRecorder) Validate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockAuthService)(nil).Validate), arg0, arg1)
}
