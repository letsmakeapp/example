// Code generated by MockGen. DO NOT EDIT.
// Source: example/abstract (interfaces: StringLengthCounter)

// Package abstract_mock is a generated GoMock package.
package abstract_mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStringLengthCounter is a mock of StringLengthCounter interface.
type MockStringLengthCounter struct {
	ctrl     *gomock.Controller
	recorder *MockStringLengthCounterMockRecorder
}

// MockStringLengthCounterMockRecorder is the mock recorder for MockStringLengthCounter.
type MockStringLengthCounterMockRecorder struct {
	mock *MockStringLengthCounter
}

// NewMockStringLengthCounter creates a new mock instance.
func NewMockStringLengthCounter(ctrl *gomock.Controller) *MockStringLengthCounter {
	mock := &MockStringLengthCounter{ctrl: ctrl}
	mock.recorder = &MockStringLengthCounterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStringLengthCounter) EXPECT() *MockStringLengthCounterMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockStringLengthCounter) Count(arg0 string) uint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", arg0)
	ret0, _ := ret[0].(uint)
	return ret0
}

// Count indicates an expected call of Count.
func (mr *MockStringLengthCounterMockRecorder) Count(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockStringLengthCounter)(nil).Count), arg0)
}
