// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/reaction-eng/restlib/email (interfaces: SmtpConnection,Mail)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	email "github.com/reaction-eng/restlib/email"
	io "io"
	reflect "reflect"
)

// MockSmtpConnection is a mock of SmtpConnection interface
type MockSmtpConnection struct {
	ctrl     *gomock.Controller
	recorder *MockSmtpConnectionMockRecorder
}

// MockSmtpConnectionMockRecorder is the mock recorder for MockSmtpConnection
type MockSmtpConnectionMockRecorder struct {
	mock *MockSmtpConnection
}

// NewMockSmtpConnection creates a new mock instance
func NewMockSmtpConnection(ctrl *gomock.Controller) *MockSmtpConnection {
	mock := &MockSmtpConnection{ctrl: ctrl}
	mock.recorder = &MockSmtpConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSmtpConnection) EXPECT() *MockSmtpConnectionMockRecorder {
	return m.recorder
}

// New mocks base method
func (m *MockSmtpConnection) New(arg0, arg1, arg2, arg3, arg4 string) email.Mail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(email.Mail)
	return ret0
}

// New indicates an expected call of New
func (mr *MockSmtpConnectionMockRecorder) New(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockSmtpConnection)(nil).New), arg0, arg1, arg2, arg3, arg4)
}

// MockMail is a mock of Mail interface
type MockMail struct {
	ctrl     *gomock.Controller
	recorder *MockMailMockRecorder
}

// MockMailMockRecorder is the mock recorder for MockMail
type MockMailMockRecorder struct {
	mock *MockMail
}

// NewMockMail creates a new mock instance
func NewMockMail(ctrl *gomock.Controller) *MockMail {
	mock := &MockMail{ctrl: ctrl}
	mock.recorder = &MockMailMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMail) EXPECT() *MockMailMockRecorder {
	return m.recorder
}

// Attach mocks base method
func (m *MockMail) Attach(arg0 string, arg1 io.Reader) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Attach", arg0, arg1)
}

// Attach indicates an expected call of Attach
func (mr *MockMailMockRecorder) Attach(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attach", reflect.TypeOf((*MockMail)(nil).Attach), arg0, arg1)
}

// Bcc mocks base method
func (m *MockMail) Bcc(arg0 ...string) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Bcc", varargs...)
}

// Bcc indicates an expected call of Bcc
func (mr *MockMailMockRecorder) Bcc(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bcc", reflect.TypeOf((*MockMail)(nil).Bcc), arg0...)
}

// From mocks base method
func (m *MockMail) From(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "From", arg0)
}

// From indicates an expected call of From
func (mr *MockMailMockRecorder) From(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "From", reflect.TypeOf((*MockMail)(nil).From), arg0)
}

// Html mocks base method
func (m *MockMail) Html() io.Writer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Html")
	ret0, _ := ret[0].(io.Writer)
	return ret0
}

// Html indicates an expected call of Html
func (mr *MockMailMockRecorder) Html() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Html", reflect.TypeOf((*MockMail)(nil).Html))
}

// ReplyTo mocks base method
func (m *MockMail) ReplyTo(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReplyTo", arg0)
}

// ReplyTo indicates an expected call of ReplyTo
func (mr *MockMailMockRecorder) ReplyTo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplyTo", reflect.TypeOf((*MockMail)(nil).ReplyTo), arg0)
}

// Send mocks base method
func (m *MockMail) Send() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send")
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockMailMockRecorder) Send() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockMail)(nil).Send))
}

// SetPlain mocks base method
func (m *MockMail) SetPlain(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPlain", arg0)
}

// SetPlain indicates an expected call of SetPlain
func (mr *MockMailMockRecorder) SetPlain(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPlain", reflect.TypeOf((*MockMail)(nil).SetPlain), arg0)
}

// Subject mocks base method
func (m *MockMail) Subject(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Subject", arg0)
}

// Subject indicates an expected call of Subject
func (mr *MockMailMockRecorder) Subject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subject", reflect.TypeOf((*MockMail)(nil).Subject), arg0)
}

// To mocks base method
func (m *MockMail) To(arg0 ...string) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "To", varargs...)
}

// To indicates an expected call of To
func (mr *MockMailMockRecorder) To(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "To", reflect.TypeOf((*MockMail)(nil).To), arg0...)
}
