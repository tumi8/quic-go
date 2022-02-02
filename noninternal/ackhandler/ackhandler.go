package ackhandler

import (
	"github.com/zirngibl/quic-go/noninternal/protocol"
	"github.com/zirngibl/quic-go/noninternal/utils"
	"github.com/zirngibl/quic-go/logging"
)

// NewAckHandler creates a new SentPacketHandler and a new ReceivedPacketHandler
func NewAckHandler(
	initialPacketNumber protocol.PacketNumber,
	initialMaxDatagramSize protocol.ByteCount,
	rttStats *utils.RTTStats,
	pers protocol.Perspective,
	tracer logging.ConnectionTracer,
	logger utils.Logger,
	version protocol.VersionNumber,
) (SentPacketHandler, ReceivedPacketHandler) {
	sph := newSentPacketHandler(initialPacketNumber, initialMaxDatagramSize, rttStats, pers, tracer, logger)
	return sph, newReceivedPacketHandler(sph, rttStats, logger, version)
}
