package ackhandler

import "github.com/tumi8/quic-go/noninternal/wire"

// FrameHandler handles the acknowledgement and the loss of a frame.
type FrameHandler interface {
	OnAcked(wire.Frame)
	OnLost(wire.Frame)
}

type Frame struct {
	Frame   wire.Frame // nil if the frame has already been acknowledged in another packet
	Handler FrameHandler
}

type StreamFrame struct {
	Frame   *wire.StreamFrame
	Handler FrameHandler
}
