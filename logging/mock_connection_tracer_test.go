// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/zirngibl/quic-go/logging (interfaces: ConnectionTracer)

// Package logging is a generated GoMock package.
package logging

import (
	net "net"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/zirngibl/quic-go/noninternal/protocol"
	utils "github.com/zirngibl/quic-go/noninternal/utils"
	wire "github.com/zirngibl/quic-go/noninternal/wire"
)

// MockConnectionTracer is a mock of ConnectionTracer interface.
type MockConnectionTracer struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionTracerMockRecorder
}

// MockConnectionTracerMockRecorder is the mock recorder for MockConnectionTracer.
type MockConnectionTracerMockRecorder struct {
	mock *MockConnectionTracer
}

// NewMockConnectionTracer creates a new mock instance.
func NewMockConnectionTracer(ctrl *gomock.Controller) *MockConnectionTracer {
	mock := &MockConnectionTracer{ctrl: ctrl}
	mock.recorder = &MockConnectionTracerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectionTracer) EXPECT() *MockConnectionTracerMockRecorder {
	return m.recorder
}

// AcknowledgedPacket mocks base method.
func (m *MockConnectionTracer) AcknowledgedPacket(arg0 protocol.EncryptionLevel, arg1 protocol.PacketNumber) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AcknowledgedPacket", arg0, arg1)
}

// AcknowledgedPacket indicates an expected call of AcknowledgedPacket.
func (mr *MockConnectionTracerMockRecorder) AcknowledgedPacket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcknowledgedPacket", reflect.TypeOf((*MockConnectionTracer)(nil).AcknowledgedPacket), arg0, arg1)
}

// BufferedPacket mocks base method.
func (m *MockConnectionTracer) BufferedPacket(arg0 PacketType) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BufferedPacket", arg0)
}

// BufferedPacket indicates an expected call of BufferedPacket.
func (mr *MockConnectionTracerMockRecorder) BufferedPacket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BufferedPacket", reflect.TypeOf((*MockConnectionTracer)(nil).BufferedPacket), arg0)
}

// Close mocks base method.
func (m *MockConnectionTracer) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockConnectionTracerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConnectionTracer)(nil).Close))
}

// ClosedConnection mocks base method.
func (m *MockConnectionTracer) ClosedConnection(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClosedConnection", arg0)
}

// ClosedConnection indicates an expected call of ClosedConnection.
func (mr *MockConnectionTracerMockRecorder) ClosedConnection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClosedConnection", reflect.TypeOf((*MockConnectionTracer)(nil).ClosedConnection), arg0)
}

// Debug mocks base method.
func (m *MockConnectionTracer) Debug(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Debug", arg0, arg1)
}

// Debug indicates an expected call of Debug.
func (mr *MockConnectionTracerMockRecorder) Debug(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockConnectionTracer)(nil).Debug), arg0, arg1)
}

// DroppedEncryptionLevel mocks base method.
func (m *MockConnectionTracer) DroppedEncryptionLevel(arg0 protocol.EncryptionLevel) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DroppedEncryptionLevel", arg0)
}

// DroppedEncryptionLevel indicates an expected call of DroppedEncryptionLevel.
func (mr *MockConnectionTracerMockRecorder) DroppedEncryptionLevel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DroppedEncryptionLevel", reflect.TypeOf((*MockConnectionTracer)(nil).DroppedEncryptionLevel), arg0)
}

// DroppedKey mocks base method.
func (m *MockConnectionTracer) DroppedKey(arg0 protocol.KeyPhase) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DroppedKey", arg0)
}

// DroppedKey indicates an expected call of DroppedKey.
func (mr *MockConnectionTracerMockRecorder) DroppedKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DroppedKey", reflect.TypeOf((*MockConnectionTracer)(nil).DroppedKey), arg0)
}

// DroppedPacket mocks base method.
func (m *MockConnectionTracer) DroppedPacket(arg0 PacketType, arg1 protocol.ByteCount, arg2 PacketDropReason) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DroppedPacket", arg0, arg1, arg2)
}

// DroppedPacket indicates an expected call of DroppedPacket.
func (mr *MockConnectionTracerMockRecorder) DroppedPacket(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DroppedPacket", reflect.TypeOf((*MockConnectionTracer)(nil).DroppedPacket), arg0, arg1, arg2)
}

// LossTimerCanceled mocks base method.
func (m *MockConnectionTracer) LossTimerCanceled() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LossTimerCanceled")
}

// LossTimerCanceled indicates an expected call of LossTimerCanceled.
func (mr *MockConnectionTracerMockRecorder) LossTimerCanceled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LossTimerCanceled", reflect.TypeOf((*MockConnectionTracer)(nil).LossTimerCanceled))
}

// LossTimerExpired mocks base method.
func (m *MockConnectionTracer) LossTimerExpired(arg0 TimerType, arg1 protocol.EncryptionLevel) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LossTimerExpired", arg0, arg1)
}

// LossTimerExpired indicates an expected call of LossTimerExpired.
func (mr *MockConnectionTracerMockRecorder) LossTimerExpired(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LossTimerExpired", reflect.TypeOf((*MockConnectionTracer)(nil).LossTimerExpired), arg0, arg1)
}

// LostPacket mocks base method.
func (m *MockConnectionTracer) LostPacket(arg0 protocol.EncryptionLevel, arg1 protocol.PacketNumber, arg2 PacketLossReason) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LostPacket", arg0, arg1, arg2)
}

// LostPacket indicates an expected call of LostPacket.
func (mr *MockConnectionTracerMockRecorder) LostPacket(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LostPacket", reflect.TypeOf((*MockConnectionTracer)(nil).LostPacket), arg0, arg1, arg2)
}

// NegotiatedVersion mocks base method.
func (m *MockConnectionTracer) NegotiatedVersion(arg0 protocol.VersionNumber, arg1, arg2 []protocol.VersionNumber) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NegotiatedVersion", arg0, arg1, arg2)
}

// NegotiatedVersion indicates an expected call of NegotiatedVersion.
func (mr *MockConnectionTracerMockRecorder) NegotiatedVersion(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NegotiatedVersion", reflect.TypeOf((*MockConnectionTracer)(nil).NegotiatedVersion), arg0, arg1, arg2)
}

// ReceivedPacket mocks base method.
func (m *MockConnectionTracer) ReceivedPacket(arg0 *wire.ExtendedHeader, arg1 protocol.ByteCount, arg2 []Frame) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReceivedPacket", arg0, arg1, arg2)
}

// ReceivedPacket indicates an expected call of ReceivedPacket.
func (mr *MockConnectionTracerMockRecorder) ReceivedPacket(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceivedPacket", reflect.TypeOf((*MockConnectionTracer)(nil).ReceivedPacket), arg0, arg1, arg2)
}

// ReceivedRetry mocks base method.
func (m *MockConnectionTracer) ReceivedRetry(arg0 *wire.Header) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReceivedRetry", arg0)
}

// ReceivedRetry indicates an expected call of ReceivedRetry.
func (mr *MockConnectionTracerMockRecorder) ReceivedRetry(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceivedRetry", reflect.TypeOf((*MockConnectionTracer)(nil).ReceivedRetry), arg0)
}

// ReceivedTransportParameters mocks base method.
func (m *MockConnectionTracer) ReceivedTransportParameters(arg0 *wire.TransportParameters) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReceivedTransportParameters", arg0)
}

// ReceivedTransportParameters indicates an expected call of ReceivedTransportParameters.
func (mr *MockConnectionTracerMockRecorder) ReceivedTransportParameters(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceivedTransportParameters", reflect.TypeOf((*MockConnectionTracer)(nil).ReceivedTransportParameters), arg0)
}

// ReceivedVersionNegotiationPacket mocks base method.
func (m *MockConnectionTracer) ReceivedVersionNegotiationPacket(arg0 *wire.Header, arg1 []protocol.VersionNumber) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReceivedVersionNegotiationPacket", arg0, arg1)
}

// ReceivedVersionNegotiationPacket indicates an expected call of ReceivedVersionNegotiationPacket.
func (mr *MockConnectionTracerMockRecorder) ReceivedVersionNegotiationPacket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceivedVersionNegotiationPacket", reflect.TypeOf((*MockConnectionTracer)(nil).ReceivedVersionNegotiationPacket), arg0, arg1)
}

// RestoredTransportParameters mocks base method.
func (m *MockConnectionTracer) RestoredTransportParameters(arg0 *wire.TransportParameters) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RestoredTransportParameters", arg0)
}

// RestoredTransportParameters indicates an expected call of RestoredTransportParameters.
func (mr *MockConnectionTracerMockRecorder) RestoredTransportParameters(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoredTransportParameters", reflect.TypeOf((*MockConnectionTracer)(nil).RestoredTransportParameters), arg0)
}

// SentPacket mocks base method.
func (m *MockConnectionTracer) SentPacket(arg0 *wire.ExtendedHeader, arg1 protocol.ByteCount, arg2 *wire.AckFrame, arg3 []Frame) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SentPacket", arg0, arg1, arg2, arg3)
}

// SentPacket indicates an expected call of SentPacket.
func (mr *MockConnectionTracerMockRecorder) SentPacket(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SentPacket", reflect.TypeOf((*MockConnectionTracer)(nil).SentPacket), arg0, arg1, arg2, arg3)
}

// SentTransportParameters mocks base method.
func (m *MockConnectionTracer) SentTransportParameters(arg0 *wire.TransportParameters) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SentTransportParameters", arg0)
}

// SentTransportParameters indicates an expected call of SentTransportParameters.
func (mr *MockConnectionTracerMockRecorder) SentTransportParameters(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SentTransportParameters", reflect.TypeOf((*MockConnectionTracer)(nil).SentTransportParameters), arg0)
}

// SetLossTimer mocks base method.
func (m *MockConnectionTracer) SetLossTimer(arg0 TimerType, arg1 protocol.EncryptionLevel, arg2 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLossTimer", arg0, arg1, arg2)
}

// SetLossTimer indicates an expected call of SetLossTimer.
func (mr *MockConnectionTracerMockRecorder) SetLossTimer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLossTimer", reflect.TypeOf((*MockConnectionTracer)(nil).SetLossTimer), arg0, arg1, arg2)
}

// StartedConnection mocks base method.
func (m *MockConnectionTracer) StartedConnection(arg0, arg1 net.Addr, arg2, arg3 protocol.ConnectionID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartedConnection", arg0, arg1, arg2, arg3)
}

// StartedConnection indicates an expected call of StartedConnection.
func (mr *MockConnectionTracerMockRecorder) StartedConnection(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartedConnection", reflect.TypeOf((*MockConnectionTracer)(nil).StartedConnection), arg0, arg1, arg2, arg3)
}

// UpdatedCongestionState mocks base method.
func (m *MockConnectionTracer) UpdatedCongestionState(arg0 CongestionState) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatedCongestionState", arg0)
}

// UpdatedCongestionState indicates an expected call of UpdatedCongestionState.
func (mr *MockConnectionTracerMockRecorder) UpdatedCongestionState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatedCongestionState", reflect.TypeOf((*MockConnectionTracer)(nil).UpdatedCongestionState), arg0)
}

// UpdatedKey mocks base method.
func (m *MockConnectionTracer) UpdatedKey(arg0 protocol.KeyPhase, arg1 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatedKey", arg0, arg1)
}

// UpdatedKey indicates an expected call of UpdatedKey.
func (mr *MockConnectionTracerMockRecorder) UpdatedKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatedKey", reflect.TypeOf((*MockConnectionTracer)(nil).UpdatedKey), arg0, arg1)
}

// UpdatedKeyFromTLS mocks base method.
func (m *MockConnectionTracer) UpdatedKeyFromTLS(arg0 protocol.EncryptionLevel, arg1 protocol.Perspective) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatedKeyFromTLS", arg0, arg1)
}

// UpdatedKeyFromTLS indicates an expected call of UpdatedKeyFromTLS.
func (mr *MockConnectionTracerMockRecorder) UpdatedKeyFromTLS(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatedKeyFromTLS", reflect.TypeOf((*MockConnectionTracer)(nil).UpdatedKeyFromTLS), arg0, arg1)
}

// UpdatedMetrics mocks base method.
func (m *MockConnectionTracer) UpdatedMetrics(arg0 *utils.RTTStats, arg1, arg2 protocol.ByteCount, arg3 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatedMetrics", arg0, arg1, arg2, arg3)
}

// UpdatedMetrics indicates an expected call of UpdatedMetrics.
func (mr *MockConnectionTracerMockRecorder) UpdatedMetrics(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatedMetrics", reflect.TypeOf((*MockConnectionTracer)(nil).UpdatedMetrics), arg0, arg1, arg2, arg3)
}

// UpdatedPTOCount mocks base method.
func (m *MockConnectionTracer) UpdatedPTOCount(arg0 uint32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatedPTOCount", arg0)
}

// UpdatedPTOCount indicates an expected call of UpdatedPTOCount.
func (mr *MockConnectionTracerMockRecorder) UpdatedPTOCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatedPTOCount", reflect.TypeOf((*MockConnectionTracer)(nil).UpdatedPTOCount), arg0)
}
