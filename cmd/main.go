package main

import (
	"log"
	"net"

	"specmatic-order-bff-grpc-go/internal/com/store/order/bff/handlers"
	"specmatic-order-bff-grpc-go/internal/com/store/order/bff/services"
	"specmatic-order-bff-grpc-go/internal/com/store/order/bff/utils"

	bff_pb "specmatic-order-bff-grpc-go/pkg/api/in/specmatic/examples/store/order_bff_grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Service addresses can be loaded from configuration (e.g., YAML or environment variables)
	orderServiceAddress := "localhost:9000"
	productServiceAddress := "localhost:9000"

	// Connect to domain services
	orderConn, err := utils.ConnectToService(orderServiceAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer orderConn.Close()

	productConn, err := utils.ConnectToService(productServiceAddress)
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
