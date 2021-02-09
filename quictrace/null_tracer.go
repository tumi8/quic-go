// +build !quictrace

package quictrace

import "gitlab.lrz.de/netintum/projects/gino/students/quic-go/noninternal/protocol"

// NewTracer returns a new Tracer that doesn't do anything.
func NewTracer() Tracer {
	return &nullTracer{}
}

type nullTracer struct{}

var _ Tracer = &nullTracer{}

func (t *nullTracer) Trace(protocol.ConnectionID, Event) {}
func (t *nullTracer) GetAllTraces() map[string][]byte    { return make(map[string][]byte) }
