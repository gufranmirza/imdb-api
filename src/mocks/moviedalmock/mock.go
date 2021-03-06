// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gufranmirza/imdb-api/dal/moviedal (interfaces: MovieDal)

// Package moviedalmock is a generated GoMock package.
package moviedalmock

import (
	gomock "github.com/golang/mock/gomock"
	dbmodels "github.com/gufranmirza/imdb-api/database/dbmodels"
	reflect "reflect"
)

// MockMovieDal is a mock of MovieDal interface
type MockMovieDal struct {
	ctrl     *gomock.Controller
	recorder *MockMovieDalMockRecorder
}

// MockMovieDalMockRecorder is the mock recorder for MockMovieDal
type MockMovieDalMockRecorder struct {
	mock *MockMovieDal
}

// NewMockMovieDal creates a new mock instance
func NewMockMovieDal(ctrl *gomock.Controller) *MockMovieDal {
	mock := &MockMovieDal{ctrl: ctrl}
	mock.recorder = &MockMovieDalMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMovieDal) EXPECT() *MockMovieDalMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockMovieDal) Create(arg0 string, arg1 *dbmodels.Movie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockMovieDalMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMovieDal)(nil).Create), arg0, arg1)
}

// Search mocks base method
func (m *MockMovieDal) Search(arg0 string) ([]dbmodels.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", arg0)
	ret0, _ := ret[0].([]dbmodels.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockMovieDalMockRecorder) Search(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockMovieDal)(nil).Search), arg0)
}
