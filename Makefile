.PHONY: all updateProto grpc grpcGateway setupTools clean

all: setupTools updateProto

updateProto: grpc grpcGateway

grpc:
	rm -rf meet/java/io/channel/api/proto
	mkdir -p meet/java/io/channel/api/proto
	protoc -I. \
		-I${GOPATH}/pkg/mod/github.com/googleapis/googleapis@v0.0.0-20220201063650-f78745822aad \
		--plugin=protoc-gen-grpc-java=/Users/max/desktop/grpc-java/compiler/build/exe/java_plugin/protoc-gen-grpc-java \
		--go_opt=module=github.com/channel-io/ch-proto \
		--go_out=. \
		--go-grpc_opt=module=github.com/channel-io/ch-proto \
		--go-grpc_out=. \
		--java_out=meet/java \
		--grpc-java_out=meet/java \
		meet/meet.proto
	protoc -I. \
		-I${GOPATH}/pkg/mod/github.com/googleapis/googleapis@v0.0.0-20220201063650-f78745822aad \
		-I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.7.3 \
		--go_opt=module=github.com/channel-io/ch-proto \
		--go_out=. \
		--go-grpc_opt=module=github.com/channel-io/ch-proto \
		--go-grpc_out=. \
		sip/sip.proto
# Setup & Install tools
setupTools:
	go mod download
	cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install % | echo "Install executables"

clean:
	go mod tidy
