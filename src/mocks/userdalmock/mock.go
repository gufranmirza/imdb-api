// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gufranmirza/imdb-api/dal/userdal (interfaces: UserDal)

// Package userdalmock is a generated GoMock package.
package userdalmock

import (
	gomock "github.com/golang/mock/gomock"
	dbmodels "github.com/gufranmirza/imdb-api/database/dbmodels"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	reflect "reflect"
)

// MockUserDal is a mock of UserDal interface
type MockUserDal struct {
	ctrl     *gomock.Controller
	recorder *MockUserDalMockRecorder
}

// MockUserDalMockRecorder is the mock recorder for MockUserDal
type MockUserDalMockRecorder struct {
	mock *MockUserDal
}

// NewMockUserDal creates a new mock instance
func NewMockUserDal(ctrl *gomock.Controller) *MockUserDal {
	mock := &MockUserDal{ctrl: ctrl}
	mock.recorder = &MockUserDalMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserDal) EXPECT() *MockUserDalMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockUserDal) Create(arg0 string, arg1 *dbmodels.User) (*dbmodels.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*dbmodels.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockUserDalMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserDal)(nil).Create), arg0, arg1)
}

// GetByEmail mocks base method
func (m *MockUserDal) GetByEmail(arg0 string) (*dbmodels.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", arg0)
	ret0, _ := ret[0].(*dbmodels.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail
func (mr *MockUserDalMockRecorder) GetByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockUserDal)(nil).GetByEmail), arg0)
}

// GetByID mocks base method
func (m *MockUserDal) GetByID(arg0 primitive.ObjectID) (*dbmodels.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*dbmodels.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockUserDalMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUserDal)(nil).GetByID), arg0)
}

// Update mocks base method
func (m *MockUserDal) Update(arg0 *dbmodels.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockUserDalMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserDal)(nil).Update), arg0)
}
