package main

import (
	"log"
	"net"
	"os"

	"specmatic-order-bff-grpc-go/internal/com/store/order/bff/handlers"
	"specmatic-order-bff-grpc-go/internal/com/store/order/bff/services"
	"specmatic-order-bff-grpc-go/internal/com/store/order/bff/utils"

	bff_pb "specmatic-order-bff-grpc-go/pkg/api/in/specmatic/examples/store/order_bff_grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	domainServerPort := os.Getenv("DOMAIN_SERVER_PORT")
	// Service addresses can be loaded from configuration (e.g., YAML or environment variables)
	orderServiceAddress := "host.docker.internal:" + domainServerPort
	productServiceAddress := "host.docker.internal:" + domainServerPort

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

	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting gRPC server on :8090")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
