module test

go 1.22.4

toolchain go1.23.2

// The version doesn't matter here, as we're replacing it with the currently checked out code anyway.
require github.com/tumi8/quic-go v0.21.0

require (
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/google/pprof v0.0.0-20240424215950-a892ee059fd6 // indirect
	github.com/onsi/ginkgo/v2 v2.19.0 // indirect
	github.com/quic-go/qpack v0.4.0 // indirect
	go.uber.org/mock v0.4.0 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	golang.org/x/exp v0.0.0-20240613232115-7f521ea00fb8 // indirect
	golang.org/x/mod v0.18.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	golang.org/x/tools v0.22.0 // indirect
)

replace github.com/tumi8/quic-go => ../../
