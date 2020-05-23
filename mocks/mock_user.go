// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/reaction-eng/restlib/users (interfaces: User)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUser is a mock of User interface
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// Activated mocks base method
func (m *MockUser) Activated() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Activated")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Activated indicates an expected call of Activated
func (mr *MockUserMockRecorder) Activated() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Activated", reflect.TypeOf((*MockUser)(nil).Activated))
}

// Email mocks base method
func (m *MockUser) Email() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Email")
	ret0, _ := ret[0].(string)
	return ret0
}

// Email indicates an expected call of Email
func (mr *MockUserMockRecorder) Email() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Email", reflect.TypeOf((*MockUser)(nil).Email))
}

// Id mocks base method
func (m *MockUser) Id() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(int)
	return ret0
}

// Id indicates an expected call of Id
func (mr *MockUserMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockUser)(nil).Id))
}

// Organizations mocks base method
func (m *MockUser) Organizations() []int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Organizations")
	ret0, _ := ret[0].([]int)
	return ret0
}

// Organizations indicates an expected call of Organizations
func (mr *MockUserMockRecorder) Organizations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Organizations", reflect.TypeOf((*MockUser)(nil).Organizations))
}

// Password mocks base method
func (m *MockUser) Password() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Password")
	ret0, _ := ret[0].(string)
	return ret0
}

// Password indicates an expected call of Password
func (mr *MockUserMockRecorder) Password() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Password", reflect.TypeOf((*MockUser)(nil).Password))
}

// PasswordLogin mocks base method
func (m *MockUser) PasswordLogin() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PasswordLogin")
	ret0, _ := ret[0].(bool)
	return ret0
}

// PasswordLogin indicates an expected call of PasswordLogin
func (mr *MockUserMockRecorder) PasswordLogin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PasswordLogin", reflect.TypeOf((*MockUser)(nil).PasswordLogin))
}

// SetEmail mocks base method
func (m *MockUser) SetEmail(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetEmail", arg0)
}

// SetEmail indicates an expected call of SetEmail
func (mr *MockUserMockRecorder) SetEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEmail", reflect.TypeOf((*MockUser)(nil).SetEmail), arg0)
}

// SetId mocks base method
func (m *MockUser) SetId(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetId", arg0)
}

// SetId indicates an expected call of SetId
func (mr *MockUserMockRecorder) SetId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetId", reflect.TypeOf((*MockUser)(nil).SetId), arg0)
}

// SetOrganizations mocks base method
func (m *MockUser) SetOrganizations(arg0 ...int) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "SetOrganizations", varargs...)
}

// SetOrganizations indicates an expected call of SetOrganizations
func (mr *MockUserMockRecorder) SetOrganizations(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOrganizations", reflect.TypeOf((*MockUser)(nil).SetOrganizations), arg0...)
}

// SetPassword mocks base method
func (m *MockUser) SetPassword(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPassword", arg0)
}

// SetPassword indicates an expected call of SetPassword
func (mr *MockUserMockRecorder) SetPassword(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPassword", reflect.TypeOf((*MockUser)(nil).SetPassword), arg0)
}

// SetToken mocks base method
func (m *MockUser) SetToken(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetToken", arg0)
}

// SetToken indicates an expected call of SetToken
func (mr *MockUserMockRecorder) SetToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetToken", reflect.TypeOf((*MockUser)(nil).SetToken), arg0)
}

// Token mocks base method
func (m *MockUser) Token() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token")
	ret0, _ := ret[0].(string)
	return ret0
}

// Token indicates an expected call of Token
func (mr *MockUserMockRecorder) Token() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockUser)(nil).Token))
}