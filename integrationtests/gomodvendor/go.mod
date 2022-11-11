module test

go 1.16

require (
	github.com/cheekybits/genny v1.0.0 // indirect
	// The version doesn't matter here, as we're replacing it with the currently checked out code anyway.
	github.com/lucas-clemente/quic-go v0.28.0
	github.com/marten-seemann/qtls-go1-16 v0.1.5 // indirect
	github.com/marten-seemann/qtls-go1-17 v0.1.2 // indirect
	github.com/marten-seemann/qtls-go1-18 v0.1.2 // indirect
	github.com/marten-seemann/qtls-go1-19 v0.1.0-beta.1 // indirect
)

replace github.com/lucas-clemente/quic-go => ../../
