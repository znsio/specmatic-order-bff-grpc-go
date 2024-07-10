package utils

import (
	"fmt"
	"google.golang.org/grpc"
)

func ConnectToService(address string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to service at %s: %v", address, err)
	}
	return conn, nil
}
