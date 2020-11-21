// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gufranmirza/imdb-api/web/server (interfaces: Server)

// Package servermock is a generated GoMock package.
package servermock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockServer is a mock of Server interface
type MockServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerMockRecorder
}

// MockServerMockRecorder is the mock recorder for MockServer
type MockServerMockRecorder struct {
	mock *MockServer
}

// NewMockServer creates a new mock instance
func NewMockServer(ctrl *gomock.Controller) *MockServer {
	mock := &MockServer{ctrl: ctrl}
	mock.recorder = &MockServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServer) EXPECT() *MockServerMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockServer) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start
func (mr *MockServerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockServer)(nil).Start))
}

// StartTimestampUTC mocks base method
func (m *MockServer) StartTimestampUTC() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartTimestampUTC")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// StartTimestampUTC indicates an expected call of StartTimestampUTC
func (mr *MockServerMockRecorder) StartTimestampUTC() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartTimestampUTC", reflect.TypeOf((*MockServer)(nil).StartTimestampUTC))
}