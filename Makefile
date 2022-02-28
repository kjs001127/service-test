.PHONY: all updateProto grpc grpcGateway setupTools clean

all: setupTools updateProto

updateProto: grpc grpcGateway

grpc:
	protoc -I. \
		-I${GOPATH}/pkg/mod/github.com/googleapis/googleapis@v0.0.0-20220201063650-f78745822aad \
        -I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.7.3 \
		--go_opt=module=github.com/channel-io/ch-proto \
		--go_out=. \
		--go-grpc_opt=module=github.com/channel-io/ch-proto \
		--go-grpc_out=. \
		call/call.proto

grpcGateway:
	protoc -I. \
		-I${GOPATH}/pkg/mod/github.com/googleapis/googleapis@v0.0.0-20220201063650-f78745822aad \
        -I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.7.3 \
		--grpc-gateway_out . \
        --grpc-gateway_opt logtostderr=true \
        --grpc-gateway_opt paths=source_relative \
        --grpc-gateway_opt generate_unbound_methods=true \
        call/call.proto

# Setup & Install tools
setupTools:
	go mod download
	cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install % | echo "Install executables"

clean:
	go mod tidy
