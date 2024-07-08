package tools

import (
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/bufbuild/protovalidate-go"
	_ "github.com/cespare/reflex@latest"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
