// +build debug

// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

package security

import (
	reflect "reflect"

	common "github.com/CJ-Jackson/paniko/paniko/common"
	shared "github.com/CJ-Jackson/paniko/paniko/shared"
	ctx "github.com/cjtoolkit/ctx"
	gomock "github.com/golang/mock/gomock"
)

// MockUserController is a mock of UserController interface
type MockUserController struct {
	ctrl     *gomock.Controller
	recorder *MockUserControllerMockRecorder
}

// MockUserControllerMockRecorder is the mock recorder for MockUserController
type MockUserControllerMockRecorder struct {
	mock *MockUserController
}

// NewMockUserController creates a new mock instance
func NewMockUserController(ctrl *gomock.Controller) *MockUserController {
	mock := &MockUserController{ctrl: ctrl}
	mock.recorder = &MockUserControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserController) EXPECT() *MockUserControllerMockRecorder {
	return m.recorder
}

// CheckCookie mocks base method
func (m *MockUserController) CheckCookie(context ctx.Context) shared.User {
	ret := m.ctrl.Call(m, "CheckCookie", context)
	ret0, _ := ret[0].(shared.User)
	return ret0
}

// CheckCookie indicates an expected call of CheckCookie
func (mr *MockUserControllerMockRecorder) CheckCookie(context interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCookie", reflect.TypeOf((*MockUserController)(nil).CheckCookie), context)
}

// GetDep mocks base method
func (m *MockUserController) GetDep() common.ContextHandler {
	ret := m.ctrl.Call(m, "GetDep")
	ret0, _ := ret[0].(common.ContextHandler)
	return ret0
}

// GetDep indicates an expected call of GetDep
func (mr *MockUserControllerMockRecorder) GetDep() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDep", reflect.TypeOf((*MockUserController)(nil).GetDep))
}

// Login mocks base method
func (m *MockUserController) Login(context ctx.Context, username, password string) bool {
	ret := m.ctrl.Call(m, "Login", context, username, password)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Login indicates an expected call of Login
func (mr *MockUserControllerMockRecorder) Login(context, username, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserController)(nil).Login), context, username, password)
}

// UpdateUser mocks base method
func (m *MockUserController) UpdateUser(username, password string) {
	m.ctrl.Call(m, "UpdateUser", username, password)
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockUserControllerMockRecorder) UpdateUser(username, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserController)(nil).UpdateUser), username, password)
}

// Save mocks base method
func (m *MockUserController) Save() {
	m.ctrl.Call(m, "Save")
}

// Save indicates an expected call of Save
func (mr *MockUserControllerMockRecorder) Save() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockUserController)(nil).Save))
}

// Logout mocks base method
func (m *MockUserController) Logout(context ctx.Context) {
	m.ctrl.Call(m, "Logout", context)
}

// Logout indicates an expected call of Logout
func (mr *MockUserControllerMockRecorder) Logout(context interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockUserController)(nil).Logout), context)
}
