package logging

//go:generate sh -c "mockgen -package logging -self_package gitlab.lrz.de/netintum/projects/gino/students/quic-go/logging -destination mock_connection_tracer_test.go gitlab.lrz.de/netintum/projects/gino/students/quic-go/logging ConnectionTracer && goimports -w mock_connection_tracer_test.go"
//go:generate sh -c "mockgen -package logging -self_package gitlab.lrz.de/netintum/projects/gino/students/quic-go/logging -destination mock_tracer_test.go gitlab.lrz.de/netintum/projects/gino/students/quic-go/logging Tracer && goimports -w mock_tracer_test.go"
