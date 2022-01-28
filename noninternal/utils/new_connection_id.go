package utils

import (
	"gitlab.lrz.de/netintum/projects/gino/students/quic-go/noninternal/protocol"
)

// NewConnectionID is a new connection ID
type NewConnectionID struct {
	SequenceNumber      uint64
	ConnectionID        protocol.ConnectionID
	StatelessResetToken protocol.StatelessResetToken
}
