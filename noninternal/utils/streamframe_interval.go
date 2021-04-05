package utils

import "gitlab.lrz.de/netintum/projects/gino/students/quic-go/noninternal/protocol"

// ByteInterval is an interval from one ByteCount to the other
type ByteInterval struct {
	Start protocol.ByteCount
	End   protocol.ByteCount
}
