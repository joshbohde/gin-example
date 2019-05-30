// Code generated by MockGen. DO NOT EDIT.
// Source: address.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	example "github.com/joshbohde/example"
	reflect "reflect"
)

// MockAddressService is a mock of AddressService interface
type MockAddressService struct {
	ctrl     *gomock.Controller
	recorder *MockAddressServiceMockRecorder
}

// MockAddressServiceMockRecorder is the mock recorder for MockAddressService
type MockAddressServiceMockRecorder struct {
	mock *MockAddressService
}

// NewMockAddressService creates a new mock instance
func NewMockAddressService(ctrl *gomock.Controller) *MockAddressService {
	mock := &MockAddressService{ctrl: ctrl}
	mock.recorder = &MockAddressServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAddressService) EXPECT() *MockAddressServiceMockRecorder {
	return m.recorder
}

// AddressForUserId mocks base method
func (m *MockAddressService) AddressForUserId(id int) (*example.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddressForUserId", id)
	ret0, _ := ret[0].(*example.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddressForUserId indicates an expected call of AddressForUserId
func (mr *MockAddressServiceMockRecorder) AddressForUserId(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddressForUserId", reflect.TypeOf((*MockAddressService)(nil).AddressForUserId), id)
}