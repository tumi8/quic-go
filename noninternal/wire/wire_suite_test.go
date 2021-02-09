package wire

import (
	"bytes"
	"testing"

	"gitlab.lrz.de/netintum/projects/gino/students/quic-go/noninternal/protocol"
	"gitlab.lrz.de/netintum/projects/gino/students/quic-go/noninternal/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWire(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wire Suite")
}

const (
	// a QUIC version that uses the IETF frame types
	versionIETFFrames = protocol.VersionTLS
)

func encodeVarInt(i uint64) []byte {
	b := &bytes.Buffer{}
	utils.WriteVarInt(b, i)
	return b.Bytes()
}
