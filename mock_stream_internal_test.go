// Code generated by MockGen. DO NOT EDIT.
// Source: stream.go

// Package quic is a generated GoMock package.
package quic

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	ackhandler "gitlab.lrz.de/netintum/projects/gino/students/quic-go/noninternal/ackhandler"
	protocol "gitlab.lrz.de/netintum/projects/gino/students/quic-go/noninternal/protocol"
	wire "gitlab.lrz.de/netintum/projects/gino/students/quic-go/noninternal/wire"
)

// MockStreamI is a mock of StreamI interface.
type MockStreamI struct {
	ctrl     *gomock.Controller
	recorder *MockStreamIMockRecorder
}

// MockStreamIMockRecorder is the mock recorder for MockStreamI.
type MockStreamIMockRecorder struct {
	mock *MockStreamI
}

// NewMockStreamI creates a new mock instance.
func NewMockStreamI(ctrl *gomock.Controller) *MockStreamI {
	mock := &MockStreamI{ctrl: ctrl}
	mock.recorder = &MockStreamIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamI) EXPECT() *MockStreamIMockRecorder {
	return m.recorder
}

// CancelRead mocks base method.
func (m *MockStreamI) CancelRead(arg0 ErrorCode) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelRead", arg0)
}

// CancelRead indicates an expected call of CancelRead.
func (mr *MockStreamIMockRecorder) CancelRead(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelRead", reflect.TypeOf((*MockStreamI)(nil).CancelRead), arg0)
}

// CancelWrite mocks base method.
func (m *MockStreamI) CancelWrite(arg0 ErrorCode) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelWrite", arg0)
}

// CancelWrite indicates an expected call of CancelWrite.
func (mr *MockStreamIMockRecorder) CancelWrite(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelWrite", reflect.TypeOf((*MockStreamI)(nil).CancelWrite), arg0)
}

// Close mocks base method.
func (m *MockStreamI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStreamIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStreamI)(nil).Close))
}

// Context mocks base method.
func (m *MockStreamI) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockStreamIMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockStreamI)(nil).Context))
}

// Read mocks base method.
func (m *MockStreamI) Read(p []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockStreamIMockRecorder) Read(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockStreamI)(nil).Read), p)
}

// SetDeadline mocks base method.
func (m *MockStreamI) SetDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDeadline indicates an expected call of SetDeadline.
func (mr *MockStreamIMockRecorder) SetDeadline(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeadline", reflect.TypeOf((*MockStreamI)(nil).SetDeadline), t)
}

// SetReadDeadline mocks base method.
func (m *MockStreamI) SetReadDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline.
func (mr *MockStreamIMockRecorder) SetReadDeadline(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockStreamI)(nil).SetReadDeadline), t)
}

// SetWriteDeadline mocks base method.
func (m *MockStreamI) SetWriteDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWriteDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWriteDeadline indicates an expected call of SetWriteDeadline.
func (mr *MockStreamIMockRecorder) SetWriteDeadline(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteDeadline", reflect.TypeOf((*MockStreamI)(nil).SetWriteDeadline), t)
}

// StreamID mocks base method.
func (m *MockStreamI) StreamID() StreamID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamID")
	ret0, _ := ret[0].(StreamID)
	return ret0
}

// StreamID indicates an expected call of StreamID.
func (mr *MockStreamIMockRecorder) StreamID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamID", reflect.TypeOf((*MockStreamI)(nil).StreamID))
}

// Write mocks base method.
func (m *MockStreamI) Write(p []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockStreamIMockRecorder) Write(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockStreamI)(nil).Write), p)
}

// closeForShutdown mocks base method.
func (m *MockStreamI) closeForShutdown(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "closeForShutdown", arg0)
}

// closeForShutdown indicates an expected call of closeForShutdown.
func (mr *MockStreamIMockRecorder) closeForShutdown(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "closeForShutdown", reflect.TypeOf((*MockStreamI)(nil).closeForShutdown), arg0)
}

// getWindowUpdate mocks base method.
func (m *MockStreamI) getWindowUpdate() protocol.ByteCount {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getWindowUpdate")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// getWindowUpdate indicates an expected call of getWindowUpdate.
func (mr *MockStreamIMockRecorder) getWindowUpdate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getWindowUpdate", reflect.TypeOf((*MockStreamI)(nil).getWindowUpdate))
}

// handleResetStreamFrame mocks base method.
func (m *MockStreamI) handleResetStreamFrame(arg0 *wire.ResetStreamFrame) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleResetStreamFrame", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleResetStreamFrame indicates an expected call of handleResetStreamFrame.
func (mr *MockStreamIMockRecorder) handleResetStreamFrame(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleResetStreamFrame", reflect.TypeOf((*MockStreamI)(nil).handleResetStreamFrame), arg0)
}

// handleStopSendingFrame mocks base method.
func (m *MockStreamI) handleStopSendingFrame(arg0 *wire.StopSendingFrame) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "handleStopSendingFrame", arg0)
}

// handleStopSendingFrame indicates an expected call of handleStopSendingFrame.
func (mr *MockStreamIMockRecorder) handleStopSendingFrame(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleStopSendingFrame", reflect.TypeOf((*MockStreamI)(nil).handleStopSendingFrame), arg0)
}

// handleStreamFrame mocks base method.
func (m *MockStreamI) handleStreamFrame(arg0 *wire.StreamFrame) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleStreamFrame", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleStreamFrame indicates an expected call of handleStreamFrame.
func (mr *MockStreamIMockRecorder) handleStreamFrame(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleStreamFrame", reflect.TypeOf((*MockStreamI)(nil).handleStreamFrame), arg0)
}

// hasData mocks base method.
func (m *MockStreamI) hasData() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "hasData")
	ret0, _ := ret[0].(bool)
	return ret0
}

// hasData indicates an expected call of hasData.
func (mr *MockStreamIMockRecorder) hasData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "hasData", reflect.TypeOf((*MockStreamI)(nil).hasData))
}

// popStreamFrame mocks base method.
func (m *MockStreamI) popStreamFrame(maxBytes protocol.ByteCount) (*ackhandler.Frame, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "popStreamFrame", maxBytes)
	ret0, _ := ret[0].(*ackhandler.Frame)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// popStreamFrame indicates an expected call of popStreamFrame.
func (mr *MockStreamIMockRecorder) popStreamFrame(maxBytes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "popStreamFrame", reflect.TypeOf((*MockStreamI)(nil).popStreamFrame), maxBytes)
}

// updateSendWindow mocks base method.
func (m *MockStreamI) updateSendWindow(arg0 protocol.ByteCount) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "updateSendWindow", arg0)
}

// updateSendWindow indicates an expected call of updateSendWindow.
func (mr *MockStreamIMockRecorder) updateSendWindow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "updateSendWindow", reflect.TypeOf((*MockStreamI)(nil).updateSendWindow), arg0)
}
