package utils

import (
	"encoding/json"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Status struct {
	Code    int32         `json:"code"`
	Message string        `json:"message"`
	Details []interface{} `json:"details,omitempty"`
}

type CustomError struct {
	Status *Status
}

func NewCustomError(code codes.Code, message string, details ...interface{}) error {
	return CustomError{
		Status: &Status{
			Code:    int32(code),
			Message: message,
			Details: details,
		},
	}
}

func (e CustomError) Error() string {
	return e.Status.Message
}

func (e CustomError) GRPCStatus() *status.Status {
	return status.New(codes.Code(e.Status.Code), e.Status.Message)
}

func (e CustomError) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Status        *Status     `json:"status"`
		Trailers      interface{} `json:"trailers"`
		Backtrace     []string    `json:"backtrace"`
		DetailMessage string      `json:"detailMessage"`
		Cause         interface{} `json:"cause"`
		StackTrace    interface{} `json:"stackTrace"`
	}{
		Status:        e.Status,
		Trailers:      map[string][]string{"content-type": {"application/grpc"}},
		Backtrace:     []string{},
		DetailMessage: e.Status.Message,
		Cause:         nil,
		StackTrace: map[string]interface{}{
			"depth":                90,
			"suppressedExceptions": []interface{}{},
		},
	})
}
