// Code generated by MockGen. DO NOT EDIT.
// Source: server.go

// Package quic is a generated GoMock package.
package quic

import (
	context "context"
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/zirngibl/quic-go/noninternal/protocol"
)

// MockQuicSession is a mock of QuicSession interface.
type MockQuicSession struct {
	ctrl     *gomock.Controller
	recorder *MockQuicSessionMockRecorder
}

// MockQuicSessionMockRecorder is the mock recorder for MockQuicSession.
type MockQuicSessionMockRecorder struct {
	mock *MockQuicSession
}

// NewMockQuicSession creates a new mock instance.
func NewMockQuicSession(ctrl *gomock.Controller) *MockQuicSession {
	mock := &MockQuicSession{ctrl: ctrl}
	mock.recorder = &MockQuicSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuicSession) EXPECT() *MockQuicSessionMockRecorder {
	return m.recorder
}

// AcceptStream mocks base method.
func (m *MockQuicSession) AcceptStream(arg0 context.Context) (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptStream", arg0)
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptStream indicates an expected call of AcceptStream.
func (mr *MockQuicSessionMockRecorder) AcceptStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptStream", reflect.TypeOf((*MockQuicSession)(nil).AcceptStream), arg0)
}

// AcceptUniStream mocks base method.
func (m *MockQuicSession) AcceptUniStream(arg0 context.Context) (ReceiveStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptUniStream", arg0)
	ret0, _ := ret[0].(ReceiveStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptUniStream indicates an expected call of AcceptUniStream.
func (mr *MockQuicSessionMockRecorder) AcceptUniStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptUniStream", reflect.TypeOf((*MockQuicSession)(nil).AcceptUniStream), arg0)
}

// CloseWithError mocks base method.
func (m *MockQuicSession) CloseWithError(arg0 ApplicationErrorCode, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseWithError", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseWithError indicates an expected call of CloseWithError.
func (mr *MockQuicSessionMockRecorder) CloseWithError(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseWithError", reflect.TypeOf((*MockQuicSession)(nil).CloseWithError), arg0, arg1)
}

// ConnectionState mocks base method.
func (m *MockQuicSession) ConnectionState() ConnectionState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectionState")
	ret0, _ := ret[0].(ConnectionState)
	return ret0
}

// ConnectionState indicates an expected call of ConnectionState.
func (mr *MockQuicSessionMockRecorder) ConnectionState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectionState", reflect.TypeOf((*MockQuicSession)(nil).ConnectionState))
}

// Context mocks base method.
func (m *MockQuicSession) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockQuicSessionMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockQuicSession)(nil).Context))
}

// GetVersion mocks base method.
func (m *MockQuicSession) GetVersion() protocol.VersionNumber {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersion")
	ret0, _ := ret[0].(protocol.VersionNumber)
	return ret0
}

// GetVersion indicates an expected call of GetVersion.
func (mr *MockQuicSessionMockRecorder) GetVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockQuicSession)(nil).GetVersion))
}

// HandshakeComplete mocks base method.
func (m *MockQuicSession) HandshakeComplete() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandshakeComplete")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// HandshakeComplete indicates an expected call of HandshakeComplete.
func (mr *MockQuicSessionMockRecorder) HandshakeComplete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandshakeComplete", reflect.TypeOf((*MockQuicSession)(nil).HandshakeComplete))
}

// LocalAddr mocks base method.
func (m *MockQuicSession) LocalAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// LocalAddr indicates an expected call of LocalAddr.
func (mr *MockQuicSessionMockRecorder) LocalAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddr", reflect.TypeOf((*MockQuicSession)(nil).LocalAddr))
}

// NextSession mocks base method.
func (m *MockQuicSession) NextSession() Session {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextSession")
	ret0, _ := ret[0].(Session)
	return ret0
}

// NextSession indicates an expected call of NextSession.
func (mr *MockQuicSessionMockRecorder) NextSession() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextSession", reflect.TypeOf((*MockQuicSession)(nil).NextSession))
}

// OpenStream mocks base method.
func (m *MockQuicSession) OpenStream() (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenStream")
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenStream indicates an expected call of OpenStream.
func (mr *MockQuicSessionMockRecorder) OpenStream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenStream", reflect.TypeOf((*MockQuicSession)(nil).OpenStream))
}

// OpenStreamSync mocks base method.
func (m *MockQuicSession) OpenStreamSync(arg0 context.Context) (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenStreamSync", arg0)
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenStreamSync indicates an expected call of OpenStreamSync.
func (mr *MockQuicSessionMockRecorder) OpenStreamSync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenStreamSync", reflect.TypeOf((*MockQuicSession)(nil).OpenStreamSync), arg0)
}

// OpenUniStream mocks base method.
func (m *MockQuicSession) OpenUniStream() (SendStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenUniStream")
	ret0, _ := ret[0].(SendStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenUniStream indicates an expected call of OpenUniStream.
func (mr *MockQuicSessionMockRecorder) OpenUniStream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUniStream", reflect.TypeOf((*MockQuicSession)(nil).OpenUniStream))
}

// OpenUniStreamSync mocks base method.
func (m *MockQuicSession) OpenUniStreamSync(arg0 context.Context) (SendStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenUniStreamSync", arg0)
	ret0, _ := ret[0].(SendStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenUniStreamSync indicates an expected call of OpenUniStreamSync.
func (mr *MockQuicSessionMockRecorder) OpenUniStreamSync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUniStreamSync", reflect.TypeOf((*MockQuicSession)(nil).OpenUniStreamSync), arg0)
}

// ReceiveMessage mocks base method.
func (m *MockQuicSession) ReceiveMessage() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceiveMessage")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceiveMessage indicates an expected call of ReceiveMessage.
func (mr *MockQuicSessionMockRecorder) ReceiveMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveMessage", reflect.TypeOf((*MockQuicSession)(nil).ReceiveMessage))
}

// RemoteAddr mocks base method.
func (m *MockQuicSession) RemoteAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr.
func (mr *MockQuicSessionMockRecorder) RemoteAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*MockQuicSession)(nil).RemoteAddr))
}

// SendMessage mocks base method.
func (m *MockQuicSession) SendMessage(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockQuicSessionMockRecorder) SendMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockQuicSession)(nil).SendMessage), arg0)
}

// destroy mocks base method.
func (m *MockQuicSession) destroy(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "destroy", arg0)
}

// destroy indicates an expected call of destroy.
func (mr *MockQuicSessionMockRecorder) destroy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "destroy", reflect.TypeOf((*MockQuicSession)(nil).destroy), arg0)
}

// earlySessionReady mocks base method.
func (m *MockQuicSession) earlySessionReady() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "earlySessionReady")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// earlySessionReady indicates an expected call of earlySessionReady.
func (mr *MockQuicSessionMockRecorder) earlySessionReady() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "earlySessionReady", reflect.TypeOf((*MockQuicSession)(nil).earlySessionReady))
}

// getPerspective mocks base method.
func (m *MockQuicSession) getPerspective() protocol.Perspective {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getPerspective")
	ret0, _ := ret[0].(protocol.Perspective)
	return ret0
}

// getPerspective indicates an expected call of getPerspective.
func (mr *MockQuicSessionMockRecorder) getPerspective() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getPerspective", reflect.TypeOf((*MockQuicSession)(nil).getPerspective))
}

// handlePacket mocks base method.
func (m *MockQuicSession) handlePacket(arg0 *receivedPacket) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "handlePacket", arg0)
}

// handlePacket indicates an expected call of handlePacket.
func (mr *MockQuicSessionMockRecorder) handlePacket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handlePacket", reflect.TypeOf((*MockQuicSession)(nil).handlePacket), arg0)
}

// run mocks base method.
func (m *MockQuicSession) run() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "run")
	ret0, _ := ret[0].(error)
	return ret0
}

// run indicates an expected call of run.
func (mr *MockQuicSessionMockRecorder) run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "run", reflect.TypeOf((*MockQuicSession)(nil).run))
}

// shutdown mocks base method.
func (m *MockQuicSession) shutdown() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "shutdown")
}

// shutdown indicates an expected call of shutdown.
func (mr *MockQuicSessionMockRecorder) shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "shutdown", reflect.TypeOf((*MockQuicSession)(nil).shutdown))
}
