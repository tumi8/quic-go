package utils

import "gitlab.lrz.de/netintum/projects/gino/students/quic-go/noninternal/protocol"

// PacketInterval is an interval from one PacketNumber to the other
type PacketInterval struct {
	Start protocol.PacketNumber
	End   protocol.PacketNumber
}
