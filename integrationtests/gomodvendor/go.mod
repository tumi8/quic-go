module test

go 1.16

// The version doesn't matter here, as we're replacing it with the currently checked out code anyway.
require github.com/lucas-clemente/quic-go v0.28.0

replace github.com/lucas-clemente/quic-go => ../../
