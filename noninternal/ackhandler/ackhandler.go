package ackhandler

import (
	"gitlab.lrz.de/netintum/projects/gino/students/quic-go/noninternal/protocol"
	"gitlab.lrz.de/netintum/projects/gino/students/quic-go/noninternal/utils"
	"gitlab.lrz.de/netintum/projects/gino/students/quic-go/logging"
	"gitlab.lrz.de/netintum/projects/gino/students/quic-go/quictrace"
)

// NewAckHandler creates a new SentPacketHandler and a new ReceivedPacketHandler
func NewAckHandler(
	initialPacketNumber protocol.PacketNumber,
	rttStats *utils.RTTStats,
	pers protocol.Perspective,
	traceCallback func(quictrace.Event),
	tracer logging.ConnectionTracer,
	logger utils.Logger,
	version protocol.VersionNumber,
) (SentPacketHandler, ReceivedPacketHandler) {
	sph := newSentPacketHandler(initialPacketNumber, rttStats, pers, traceCallback, tracer, logger)
	return sph, newReceivedPacketHandler(sph, rttStats, logger, version)
}
