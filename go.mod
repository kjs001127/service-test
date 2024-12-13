module github.com/channel-io/service-test

go 1.21

replace github.com/channel-io/ch-proto v0.0.0 => ./ch-proto

require (
	connectrpc.com/connect v1.17.0
	golang.org/x/net v0.23.0
	google.golang.org/protobuf v1.34.2
)

require golang.org/x/text v0.14.0 // indirect
