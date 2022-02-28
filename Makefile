
grpc.install:
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

buf.gen:
	buf generate

buf.update:
	buf mod update

buf.install:
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking@latest
	go install github.com/bufbuild/buf/cmd/protoc-gen-buf-lint@latest

go.get:
	go get -u ./...

go.gen: buf.gen

go.tidy:
	go mod tidy -compat=1.17

go.test:
	go test ./...

go.install: grpc.install buf.install
