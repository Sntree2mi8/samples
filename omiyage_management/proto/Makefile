install:
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0
	go install github.com/bufbuild/buf/cmd/buf@v1.19.0

generate:
	buf generate

lint:
	buf lint

format:
	buf format -w
