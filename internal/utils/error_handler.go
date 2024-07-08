package utils

import (
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/protobuf/proto"
)

func ValidateReq(req proto.Message) (bool, error) {
	validator, err := protovalidate.New()
	if err != nil {
		return true, err
	}
	err = validator.Validate(req)
	if err != nil {
		return true, err
	}
	return false, nil
}
