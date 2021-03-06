// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gufranmirza/imdb-api/web/services/health (interfaces: Health)

// Package healthmock is a generated GoMock package.
package healthmock

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockHealth is a mock of Health interface
type MockHealth struct {
	ctrl     *gomock.Controller
	recorder *MockHealthMockRecorder
}

// MockHealthMockRecorder is the mock recorder for MockHealth
type MockHealthMockRecorder struct {
	mock *MockHealth
}

// NewMockHealth creates a new mock instance
func NewMockHealth(ctrl *gomock.Controller) *MockHealth {
	mock := &MockHealth{ctrl: ctrl}
	mock.recorder = &MockHealthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHealth) EXPECT() *MockHealthMockRecorder {
	return m.recorder
}

// GetHealth mocks base method
func (m *MockHealth) GetHealth(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetHealth", arg0, arg1)
}

// GetHealth indicates an expected call of GetHealth
func (mr *MockHealthMockRecorder) GetHealth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHealth", reflect.TypeOf((*MockHealth)(nil).GetHealth), arg0, arg1)
}
