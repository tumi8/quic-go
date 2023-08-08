package ackhandler

import (
	"github.com/tumi8/quic-go/noninternal/protocol"
	"github.com/tumi8/quic-go/noninternal/utils"
	"github.com/tumi8/quic-go/logging"
)

// NewAckHandler creates a new SentPacketHandler and a new ReceivedPacketHandler.
// clientAddressValidated indicates whether the address was validated beforehand by an address validation token.
// clientAddressValidated has no effect for a client.
func NewAckHandler(
	initialPacketNumber protocol.PacketNumber,
	initialMaxDatagramSize protocol.ByteCount,
	rttStats *utils.RTTStats,
	clientAddressValidated bool,
	pers protocol.Perspective,
	tracer logging.ConnectionTracer,
	logger utils.Logger,
) (SentPacketHandler, ReceivedPacketHandler) {
	sph := newSentPacketHandler(initialPacketNumber, initialMaxDatagramSize, rttStats, clientAddressValidated, pers, tracer, logger)
	return sph, newReceivedPacketHandler(sph, rttStats, logger)
}
