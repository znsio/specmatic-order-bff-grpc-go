package main

import (
	"fmt"
	"log"
	"net"

	pb "specmatic-order-bff-grpc-go/pkg/api/proto_files"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Your service implementations
type orderServer struct {
	pb.UnimplementedOrderServiceServer
	// Add any fields you need
}

type productServer struct {
	pb.UnimplementedProductServiceServer
	// Add any fields you need
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Register your services
	pb.RegisterOrderServiceServer(s, &orderServer{})
	pb.RegisterProductServiceServer(s, &productServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	fmt.Println("Server is running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
