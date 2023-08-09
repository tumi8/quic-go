// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tumi8/quic-go/noninternal/handshake (interfaces: ShortHeaderOpener)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/tumi8/quic-go/noninternal/protocol"
)

// MockShortHeaderOpener is a mock of ShortHeaderOpener interface.
type MockShortHeaderOpener struct {
	ctrl     *gomock.Controller
	recorder *MockShortHeaderOpenerMockRecorder
}

// MockShortHeaderOpenerMockRecorder is the mock recorder for MockShortHeaderOpener.
type MockShortHeaderOpenerMockRecorder struct {
	mock *MockShortHeaderOpener
}

// NewMockShortHeaderOpener creates a new mock instance.
func NewMockShortHeaderOpener(ctrl *gomock.Controller) *MockShortHeaderOpener {
	mock := &MockShortHeaderOpener{ctrl: ctrl}
	mock.recorder = &MockShortHeaderOpenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShortHeaderOpener) EXPECT() *MockShortHeaderOpenerMockRecorder {
	return m.recorder
}

// DecodePacketNumber mocks base method.
func (m *MockShortHeaderOpener) DecodePacketNumber(arg0 protocol.PacketNumber, arg1 protocol.PacketNumberLen) protocol.PacketNumber {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecodePacketNumber", arg0, arg1)
	ret0, _ := ret[0].(protocol.PacketNumber)
	return ret0
}

// DecodePacketNumber indicates an expected call of DecodePacketNumber.
func (mr *MockShortHeaderOpenerMockRecorder) DecodePacketNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodePacketNumber", reflect.TypeOf((*MockShortHeaderOpener)(nil).DecodePacketNumber), arg0, arg1)
}

// DecryptHeader mocks base method.
func (m *MockShortHeaderOpener) DecryptHeader(arg0 []byte, arg1 *byte, arg2 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DecryptHeader", arg0, arg1, arg2)
}

// DecryptHeader indicates an expected call of DecryptHeader.
func (mr *MockShortHeaderOpenerMockRecorder) DecryptHeader(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecryptHeader", reflect.TypeOf((*MockShortHeaderOpener)(nil).DecryptHeader), arg0, arg1, arg2)
}

// Open mocks base method.
func (m *MockShortHeaderOpener) Open(arg0, arg1 []byte, arg2 time.Time, arg3 protocol.PacketNumber, arg4 protocol.KeyPhaseBit, arg5 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open.
func (mr *MockShortHeaderOpenerMockRecorder) Open(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockShortHeaderOpener)(nil).Open), arg0, arg1, arg2, arg3, arg4, arg5)
}