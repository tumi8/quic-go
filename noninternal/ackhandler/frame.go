package ackhandler

import "gitlab.lrz.de/netintum/projects/gino/students/quic-go/noninternal/wire"

type Frame struct {
	wire.Frame // nil if the frame has already been acknowledged in another packet
	OnLost     func(wire.Frame)
	OnAcked    func(wire.Frame)
}
