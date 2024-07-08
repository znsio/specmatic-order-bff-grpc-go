package utils

import (
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func ValidateReq(req proto.Message) error {
	validator, err := protovalidate.New()
	if err != nil {
		return NewCustomError(codes.Internal, "Failed to create validator", err.Error())
	}
	err = validator.Validate(req)
	if err != nil {
		return NewCustomError(codes.InvalidArgument, "Validation error", err.Error())
	}
	return nil
}
