package wire

import (
	"bytes"

	"gitlab.lrz.de/netintum/projects/gino/students/quic-go/noninternal/protocol"
)

// A Frame in QUIC
type Frame interface {
	Write(b *bytes.Buffer, version protocol.VersionNumber) error
	Length(version protocol.VersionNumber) protocol.ByteCount
}

// A FrameParser parses QUIC frames, one by one.
type FrameParser interface {
	ParseNext(*bytes.Reader, protocol.EncryptionLevel) (Frame, error)
	SetAckDelayExponent(uint8)
}
