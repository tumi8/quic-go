package ackhandler

//go:generate sh -c "go run github.com/golang/mock/mockgen -build_flags=\"-tags=gomock\"  -package ackhandler -destination mock_sent_packet_tracker_test.go github.com/tumi8/quic-go/noninternal/ackhandler SentPacketTracker"
type SentPacketTracker = sentPacketTracker
