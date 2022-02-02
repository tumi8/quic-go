package logging

//go:generate sh -c "mockgen -package logging -self_package github.com/zirngibl/quic-go/logging -destination mock_connection_tracer_test.go github.com/zirngibl/quic-go/logging ConnectionTracer && goimports -w mock_connection_tracer_test.go"
//go:generate sh -c "mockgen -package logging -self_package github.com/zirngibl/quic-go/logging -destination mock_tracer_test.go github.com/zirngibl/quic-go/logging Tracer && goimports -w mock_tracer_test.go"
