// Code generated by MockGen. DO NOT EDIT.
// Source: packet_packer.go

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ackhandler "gitlab.lrz.de/netintum/projects/gino/students/quic-go/noninternal/ackhandler"
	protocol "gitlab.lrz.de/netintum/projects/gino/students/quic-go/noninternal/protocol"
)

// MockFrameSource is a mock of FrameSource interface.
type MockFrameSource struct {
	ctrl     *gomock.Controller
	recorder *MockFrameSourceMockRecorder
}

// MockFrameSourceMockRecorder is the mock recorder for MockFrameSource.
type MockFrameSourceMockRecorder struct {
	mock *MockFrameSource
}

// NewMockFrameSource creates a new mock instance.
func NewMockFrameSource(ctrl *gomock.Controller) *MockFrameSource {
	mock := &MockFrameSource{ctrl: ctrl}
	mock.recorder = &MockFrameSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFrameSource) EXPECT() *MockFrameSourceMockRecorder {
	return m.recorder
}

// AppendControlFrames mocks base method.
func (m *MockFrameSource) AppendControlFrames(arg0 []ackhandler.Frame, arg1 protocol.ByteCount) ([]ackhandler.Frame, protocol.ByteCount) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendControlFrames", arg0, arg1)
	ret0, _ := ret[0].([]ackhandler.Frame)
	ret1, _ := ret[1].(protocol.ByteCount)
	return ret0, ret1
}

// AppendControlFrames indicates an expected call of AppendControlFrames.
func (mr *MockFrameSourceMockRecorder) AppendControlFrames(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendControlFrames", reflect.TypeOf((*MockFrameSource)(nil).AppendControlFrames), arg0, arg1)
}

// AppendStreamFrames mocks base method.
func (m *MockFrameSource) AppendStreamFrames(arg0 []ackhandler.Frame, arg1 protocol.ByteCount) ([]ackhandler.Frame, protocol.ByteCount) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendStreamFrames", arg0, arg1)
	ret0, _ := ret[0].([]ackhandler.Frame)
	ret1, _ := ret[1].(protocol.ByteCount)
	return ret0, ret1
}

// AppendStreamFrames indicates an expected call of AppendStreamFrames.
func (mr *MockFrameSourceMockRecorder) AppendStreamFrames(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendStreamFrames", reflect.TypeOf((*MockFrameSource)(nil).AppendStreamFrames), arg0, arg1)
}

// HasData mocks base method.
func (m *MockFrameSource) HasData() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasData")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasData indicates an expected call of HasData.
func (mr *MockFrameSourceMockRecorder) HasData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasData", reflect.TypeOf((*MockFrameSource)(nil).HasData))
}
