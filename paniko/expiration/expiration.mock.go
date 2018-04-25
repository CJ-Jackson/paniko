// +build debug

// Code generated by MockGen. DO NOT EDIT.
// Source: expiration.go

package expiration

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockExpiration is a mock of Expiration interface
type MockExpiration struct {
	ctrl     *gomock.Controller
	recorder *MockExpirationMockRecorder
}

// MockExpirationMockRecorder is the mock recorder for MockExpiration
type MockExpirationMockRecorder struct {
	mock *MockExpiration
}

// NewMockExpiration creates a new mock instance
func NewMockExpiration(ctrl *gomock.Controller) *MockExpiration {
	mock := &MockExpiration{ctrl: ctrl}
	mock.recorder = &MockExpirationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExpiration) EXPECT() *MockExpirationMockRecorder {
	return m.recorder
}

// GetTime mocks base method
func (m *MockExpiration) GetTime() time.Time {
	ret := m.ctrl.Call(m, "GetTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTime indicates an expected call of GetTime
func (mr *MockExpirationMockRecorder) GetTime() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTime", reflect.TypeOf((*MockExpiration)(nil).GetTime))
}

// Expired mocks base method
func (m *MockExpiration) Expired() bool {
	ret := m.ctrl.Call(m, "Expired")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Expired indicates an expected call of Expired
func (mr *MockExpirationMockRecorder) Expired() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Expired", reflect.TypeOf((*MockExpiration)(nil).Expired))
}

// Reset mocks base method
func (m *MockExpiration) Reset() {
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset
func (mr *MockExpirationMockRecorder) Reset() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockExpiration)(nil).Reset))
}

// Expire mocks base method
func (m *MockExpiration) Expire() {
	m.ctrl.Call(m, "Expire")
}

// Expire indicates an expected call of Expire
func (mr *MockExpirationMockRecorder) Expire() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Expire", reflect.TypeOf((*MockExpiration)(nil).Expire))
}