
buf.gen:
	buf generate

buf.update:
	buf mod update

go.get:
	go get -u ./...

go.gen: buf.gen

go.tidy:
	go mod tidy -compat=1.18

go.test:
	go test ./...

go.install:
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest

	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest

