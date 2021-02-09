package wire

import (
	"sync"

	"gitlab.lrz.de/netintum/projects/gino/students/quic-go/noninternal/protocol"
)

var pool sync.Pool

func init() {
	pool.New = func() interface{} {
		return &StreamFrame{
			Data:     make([]byte, 0, protocol.MaxReceivePacketSize),
			fromPool: true,
		}
	}
}

func GetStreamFrame() *StreamFrame {
	f := pool.Get().(*StreamFrame)
	return f
}

func putStreamFrame(f *StreamFrame) {
	if !f.fromPool {
		return
	}
	if protocol.ByteCount(cap(f.Data)) != protocol.MaxReceivePacketSize {
		panic("wire.PutStreamFrame called with packet of wrong size!")
	}
	pool.Put(f)
}
