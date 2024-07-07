package handlers

import (
	"context"
	"specmatic-order-bff-grpc-go/internal/services"
	pb "specmatic-order-bff-grpc-go/pkg/api/proto_files"
)

type OrderHandler struct {
	pb.UnimplementedOrderServiceServer
	domainAPIService *services.DomainAPIService
}

func NewOrderHandler(domainAPIService *services.DomainAPIService) *OrderHandler {
	return &OrderHandler{
		domainAPIService: domainAPIService,
	}
}

func (h *OrderHandler) FindAvailableProducts(ctx context.Context, req *pb.FindAvailableProductsRequest) (*pb.ProductListResponse, error) {
	return h.domainAPIService.FindProducts(ctx, req)
}

func (h *OrderHandler) CreateProduct(ctx context.Context, req *pb.NewProduct) (*pb.ProductId, error) {
	return h.domainAPIService.CreateProduct(ctx, req)
}

func (h *OrderHandler) CreateOrder(ctx context.Context, req *pb.NewOrder) (*pb.OrderId, error) {
	return h.domainAPIService.CreateOrder(ctx, req)
}
