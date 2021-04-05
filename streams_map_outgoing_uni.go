// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package quic

import (
	"context"
	"sync"

	"gitlab.lrz.de/netintum/projects/gino/students/quic-go/noninternal/protocol"
	"gitlab.lrz.de/netintum/projects/gino/students/quic-go/noninternal/wire"
)

type outgoingUniStreamsMap struct {
	mutex sync.RWMutex

	streams map[protocol.StreamNum]sendStreamI

	openQueue      map[uint64]chan struct{}
	lowestInQueue  uint64
	highestInQueue uint64

	nextStream  protocol.StreamNum // stream ID of the stream returned by OpenStream(Sync)
	maxStream   protocol.StreamNum // the maximum stream ID we're allowed to open
	blockedSent bool               // was a STREAMS_BLOCKED sent for the current maxStream

	newStream            func(protocol.StreamNum) sendStreamI
	queueStreamIDBlocked func(*wire.StreamsBlockedFrame)

	closeErr error
}

func newOutgoingUniStreamsMap(
	newStream func(protocol.StreamNum) sendStreamI,
	queueControlFrame func(wire.Frame),
) *outgoingUniStreamsMap {
	return &outgoingUniStreamsMap{
		streams:              make(map[protocol.StreamNum]sendStreamI),
		openQueue:            make(map[uint64]chan struct{}),
		maxStream:            protocol.InvalidStreamNum,
		nextStream:           1,
		newStream:            newStream,
		queueStreamIDBlocked: func(f *wire.StreamsBlockedFrame) { queueControlFrame(f) },
	}
}

func (m *outgoingUniStreamsMap) OpenStream() (sendStreamI, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.closeErr != nil {
		return nil, m.closeErr
	}

	// if there are OpenStreamSync calls waiting, return an error here
	if len(m.openQueue) > 0 || m.nextStream > m.maxStream {
		m.maybeSendBlockedFrame()
		return nil, streamOpenErr{errTooManyOpenStreams}
	}
	return m.openStream(), nil
}

func (m *outgoingUniStreamsMap) OpenStreamSync(ctx context.Context) (sendStreamI, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.closeErr != nil {
		return nil, m.closeErr
	}

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if len(m.openQueue) == 0 && m.nextStream <= m.maxStream {
		return m.openStream(), nil
	}

	waitChan := make(chan struct{}, 1)
	queuePos := m.highestInQueue
	m.highestInQueue++
	if len(m.openQueue) == 0 {
		m.lowestInQueue = queuePos
	}
	m.openQueue[queuePos] = waitChan
	m.maybeSendBlockedFrame()

	for {
		m.mutex.Unlock()
		select {
		case <-ctx.Done():
			m.mutex.Lock()
			delete(m.openQueue, queuePos)
			return nil, ctx.Err()
		case <-waitChan:
		}
		m.mutex.Lock()

		if m.closeErr != nil {
			return nil, m.closeErr
		}
		if m.nextStream > m.maxStream {
			// no stream available. Continue waiting
			continue
		}
		str := m.openStream()
		delete(m.openQueue, queuePos)
		m.lowestInQueue = queuePos + 1
		m.unblockOpenSync()
		return str, nil
	}
}

func (m *outgoingUniStreamsMap) openStream() sendStreamI {
	s := m.newStream(m.nextStream)
	m.streams[m.nextStream] = s
	m.nextStream++
	return s
}

// maybeSendBlockedFrame queues a STREAMS_BLOCKED frame for the current stream offset,
// if we haven't sent one for this offset yet
func (m *outgoingUniStreamsMap) maybeSendBlockedFrame() {
	if m.blockedSent {
		return
	}

	var streamNum protocol.StreamNum
	if m.maxStream != protocol.InvalidStreamNum {
		streamNum = m.maxStream
	}
	m.queueStreamIDBlocked(&wire.StreamsBlockedFrame{
		Type:        protocol.StreamTypeUni,
		StreamLimit: streamNum,
	})
	m.blockedSent = true
}

func (m *outgoingUniStreamsMap) GetStream(num protocol.StreamNum) (sendStreamI, error) {
	m.mutex.RLock()
	if num >= m.nextStream {
		m.mutex.RUnlock()
		return nil, streamError{
			message: "peer attempted to open stream %d",
			nums:    []protocol.StreamNum{num},
		}
	}
	s := m.streams[num]
	m.mutex.RUnlock()
	return s, nil
}

func (m *outgoingUniStreamsMap) DeleteStream(num protocol.StreamNum) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, ok := m.streams[num]; !ok {
		return streamError{
			message: "Tried to delete unknown outgoing stream %d",
			nums:    []protocol.StreamNum{num},
		}
	}
	delete(m.streams, num)
	return nil
}

func (m *outgoingUniStreamsMap) SetMaxStream(num protocol.StreamNum) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if num <= m.maxStream {
		return
	}
	m.maxStream = num
	m.blockedSent = false
	if m.maxStream < m.nextStream-1+protocol.StreamNum(len(m.openQueue)) {
		m.maybeSendBlockedFrame()
	}
	m.unblockOpenSync()
}

// UpdateSendWindow is called when the peer's transport parameters are received.
// Only in the case of a 0-RTT handshake will we have open streams at this point.
// We might need to update the send window, in case the server increased it.
func (m *outgoingUniStreamsMap) UpdateSendWindow(limit protocol.ByteCount) {
	m.mutex.Lock()
	for _, str := range m.streams {
		str.updateSendWindow(limit)
	}
	m.mutex.Unlock()
}

// unblockOpenSync unblocks the next OpenStreamSync go-routine to open a new stream
func (m *outgoingUniStreamsMap) unblockOpenSync() {
	if len(m.openQueue) == 0 {
		return
	}
	for qp := m.lowestInQueue; qp <= m.highestInQueue; qp++ {
		c, ok := m.openQueue[qp]
		if !ok { // entry was deleted because the context was canceled
			continue
		}
		// unblockOpenSync is called both from OpenStreamSync and from SetMaxStream.
		// It's sufficient to only unblock OpenStreamSync once.
		select {
		case c <- struct{}{}:
		default:
		}
		return
	}
}

func (m *outgoingUniStreamsMap) CloseWithError(err error) {
	m.mutex.Lock()
	m.closeErr = err
	for _, str := range m.streams {
		str.closeForShutdown(err)
	}
	for _, c := range m.openQueue {
		if c != nil {
			close(c)
		}
	}
	m.mutex.Unlock()
}
