// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	example "github.com/joshbohde/example"
	reflect "reflect"
)

// MockUserService is a mock of UserService interface
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// User mocks base method
func (m *MockUserService) User(ctx context.Context, id int) (*example.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "User", ctx, id)
	ret0, _ := ret[0].(*example.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// User indicates an expected call of User
func (mr *MockUserServiceMockRecorder) User(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "User", reflect.TypeOf((*MockUserService)(nil).User), ctx, id)
}
