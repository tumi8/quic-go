// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tumi8/quic-go/noninternal/handshake (interfaces: ShortHeaderSealer)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/tumi8/quic-go/noninternal/protocol"
)

// MockShortHeaderSealer is a mock of ShortHeaderSealer interface.
type MockShortHeaderSealer struct {
	ctrl     *gomock.Controller
	recorder *MockShortHeaderSealerMockRecorder
}

// MockShortHeaderSealerMockRecorder is the mock recorder for MockShortHeaderSealer.
type MockShortHeaderSealerMockRecorder struct {
	mock *MockShortHeaderSealer
}

// NewMockShortHeaderSealer creates a new mock instance.
func NewMockShortHeaderSealer(ctrl *gomock.Controller) *MockShortHeaderSealer {
	mock := &MockShortHeaderSealer{ctrl: ctrl}
	mock.recorder = &MockShortHeaderSealerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShortHeaderSealer) EXPECT() *MockShortHeaderSealerMockRecorder {
	return m.recorder
}

// EncryptHeader mocks base method.
func (m *MockShortHeaderSealer) EncryptHeader(arg0 []byte, arg1 *byte, arg2 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EncryptHeader", arg0, arg1, arg2)
}

// EncryptHeader indicates an expected call of EncryptHeader.
func (mr *MockShortHeaderSealerMockRecorder) EncryptHeader(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncryptHeader", reflect.TypeOf((*MockShortHeaderSealer)(nil).EncryptHeader), arg0, arg1, arg2)
}

// KeyPhase mocks base method.
func (m *MockShortHeaderSealer) KeyPhase() protocol.KeyPhaseBit {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyPhase")
	ret0, _ := ret[0].(protocol.KeyPhaseBit)
	return ret0
}

// KeyPhase indicates an expected call of KeyPhase.
func (mr *MockShortHeaderSealerMockRecorder) KeyPhase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyPhase", reflect.TypeOf((*MockShortHeaderSealer)(nil).KeyPhase))
}

// Overhead mocks base method.
func (m *MockShortHeaderSealer) Overhead() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Overhead")
	ret0, _ := ret[0].(int)
	return ret0
}

// Overhead indicates an expected call of Overhead.
func (mr *MockShortHeaderSealerMockRecorder) Overhead() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Overhead", reflect.TypeOf((*MockShortHeaderSealer)(nil).Overhead))
}

// Seal mocks base method.
func (m *MockShortHeaderSealer) Seal(arg0, arg1 []byte, arg2 protocol.PacketNumber, arg3 []byte) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seal", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Seal indicates an expected call of Seal.
func (mr *MockShortHeaderSealerMockRecorder) Seal(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seal", reflect.TypeOf((*MockShortHeaderSealer)(nil).Seal), arg0, arg1, arg2, arg3)
}
