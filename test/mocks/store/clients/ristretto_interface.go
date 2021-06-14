// Code generated by MockGen. DO NOT EDIT.
// Source: store/ristretto.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockRistrettoClientInterface is a mock of RistrettoClientInterface interface.
type MockRistrettoClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRistrettoClientInterfaceMockRecorder
}

// MockRistrettoClientInterfaceMockRecorder is the mock recorder for MockRistrettoClientInterface.
type MockRistrettoClientInterfaceMockRecorder struct {
	mock *MockRistrettoClientInterface
}

// NewMockRistrettoClientInterface creates a new mock instance.
func NewMockRistrettoClientInterface(ctrl *gomock.Controller) *MockRistrettoClientInterface {
	mock := &MockRistrettoClientInterface{ctrl: ctrl}
	mock.recorder = &MockRistrettoClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRistrettoClientInterface) EXPECT() *MockRistrettoClientInterfaceMockRecorder {
	return m.recorder
}

// Clear mocks base method.
func (m *MockRistrettoClientInterface) Clear() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Clear")
}

// Clear indicates an expected call of Clear.
func (mr *MockRistrettoClientInterfaceMockRecorder) Clear() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockRistrettoClientInterface)(nil).Clear))
}

// Del mocks base method.
func (m *MockRistrettoClientInterface) Del(key interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Del", key)
}

// Del indicates an expected call of Del.
func (mr *MockRistrettoClientInterfaceMockRecorder) Del(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockRistrettoClientInterface)(nil).Del), key)
}

// Get mocks base method.
func (m *MockRistrettoClientInterface) Get(key interface{}) (interface{}, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRistrettoClientInterfaceMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRistrettoClientInterface)(nil).Get), key)
}

// SetWithTTL mocks base method.
func (m *MockRistrettoClientInterface) SetWithTTL(key, value interface{}, cost int64, ttl time.Duration) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWithTTL", key, value, cost, ttl)
	ret0, _ := ret[0].(bool)
	return ret0
}

// SetWithTTL indicates an expected call of SetWithTTL.
func (mr *MockRistrettoClientInterfaceMockRecorder) SetWithTTL(key, value, cost, ttl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWithTTL", reflect.TypeOf((*MockRistrettoClientInterface)(nil).SetWithTTL), key, value, cost, ttl)
}
