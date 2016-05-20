// Automatically generated by MockGen. DO NOT EDIT!
// Source: output.go

package mock_gotable

import (
	gomock "github.com/golang/mock/gomock"
	. "github.com/tomcraven/gotable"
)

// Mock of Output interface
type MockOutput struct {
	ctrl     *gomock.Controller
	recorder *_MockOutputRecorder
}

// Recorder for MockOutput (not exported)
type _MockOutputRecorder struct {
	mock *MockOutput
}

func NewMockOutput(ctrl *gomock.Controller) *MockOutput {
	mock := &MockOutput{ctrl: ctrl}
	mock.recorder = &_MockOutputRecorder{mock}
	return mock
}

func (_m *MockOutput) EXPECT() *_MockOutputRecorder {
	return _m.recorder
}

func (_m *MockOutput) Print(_param0 string) Output {
	ret := _m.ctrl.Call(_m, "Print", _param0)
	ret0, _ := ret[0].(Output)
	return ret0
}

func (_mr *_MockOutputRecorder) Print(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Print", arg0)
}

func (_m *MockOutput) Flush() Output {
	ret := _m.ctrl.Call(_m, "Flush")
	ret0, _ := ret[0].(Output)
	return ret0
}

func (_mr *_MockOutputRecorder) Flush() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Flush")
}
