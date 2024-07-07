package main

import (
	"log"
	"net"

	"specmatic-order-bff-grpc-go/internal/handlers"
	"specmatic-order-bff-grpc-go/internal/services"
	pb "specmatic-order-bff-grpc-go/pkg/api/proto_files"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Set up connections to domain services
	orderConn, err := grpc.Dial("order-service-address:port", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to order service: %v", err)
	}
	defer orderConn.Close()

	productConn, err := grpc.Dial("product-service-address:port", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to product service: %v", err)
	}
	defer productConn.Close()

	domainAPIService := services.NewDomainAPIService(orderConn, productConn)

	grpcServer := grpc.NewServer()

	orderHandler := handlers.NewOrderHandler(domainAPIService)
	pb.RegisterOrderServiceServer(grpcServer, orderHandler)

	reflection.Register(grpcServer)

	log.Println("Starting gRPC server on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
