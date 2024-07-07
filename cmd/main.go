package main

import (
	"fmt"
	"log"
	"net"

	"specmatic-order-bff-grpc-go/internal/handlers"
	"specmatic-order-bff-grpc-go/internal/services"

	bff_pb "specmatic-order-bff-grpc-go/pkg/api/in/specmatic/examples/store/order_bff_grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func connectToService(address string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to service at %s: %v", address, err)
	}
	return conn, nil
}

func main() {
	// Service addresses can be loaded from configuration (e.g., YAML or environment variables)
	orderServiceAddress := "order-service-address:port"
	productServiceAddress := "product-service-address:port"

	// Connect to domain services
	orderConn, err := connectToService(orderServiceAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer orderConn.Close()

	productConn, err := connectToService(productServiceAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer productConn.Close()

	// Setup BFF gRPC server
	domainAPIService := services.NewDomainAPIService(orderConn, productConn)

	grpcServer := grpc.NewServer()

	bffHandler := handlers.NewBffHandler(domainAPIService)
	bff_pb.RegisterOrderServiceServer(grpcServer, bffHandler)

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting gRPC server on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
