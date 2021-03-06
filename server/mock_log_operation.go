// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/google/trillian/server (interfaces: LogOperation)

// Package server is a generated GoMock package.
package server

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLogOperation is a mock of LogOperation interface
type MockLogOperation struct {
	ctrl     *gomock.Controller
	recorder *MockLogOperationMockRecorder
}

// MockLogOperationMockRecorder is the mock recorder for MockLogOperation
type MockLogOperationMockRecorder struct {
	mock *MockLogOperation
}

// NewMockLogOperation creates a new mock instance
func NewMockLogOperation(ctrl *gomock.Controller) *MockLogOperation {
	mock := &MockLogOperation{ctrl: ctrl}
	mock.recorder = &MockLogOperationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogOperation) EXPECT() *MockLogOperationMockRecorder {
	return m.recorder
}

// ExecutePass mocks base method
func (m *MockLogOperation) ExecutePass(arg0 context.Context, arg1 int64, arg2 *LogOperationInfo) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecutePass", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecutePass indicates an expected call of ExecutePass
func (mr *MockLogOperationMockRecorder) ExecutePass(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecutePass", reflect.TypeOf((*MockLogOperation)(nil).ExecutePass), arg0, arg1, arg2)
}

// Name mocks base method
func (m *MockLogOperation) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockLogOperationMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockLogOperation)(nil).Name))
}
