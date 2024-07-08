package handlers

import (
	"context"
	"fmt"
	"specmatic-order-bff-grpc-go/internal/services"
	bff_pb "specmatic-order-bff-grpc-go/pkg/api/in/specmatic/examples/store/order_bff_grpc"
)

type BffHandler struct {
	bff_pb.UnimplementedOrderServiceServer
	domainAPIService *services.DomainAPIService
}

func NewBffHandler(domainAPIService *services.DomainAPIService) *BffHandler {
	return &BffHandler{
		domainAPIService: domainAPIService,
	}
}

func (h *BffHandler) FindAvailableProducts(ctx context.Context, req *bff_pb.FindAvailableProductsRequest) (*bff_pb.ProductListResponse, error) {
	if req.PageSize <= 0 {
		return nil, fmt.Errorf("PageSize must be greater than 0")
	}
	return h.domainAPIService.FindProducts(ctx, req)
}

func (h *BffHandler) CreateProduct(ctx context.Context, req *bff_pb.NewProduct) (*bff_pb.ProductId, error) {
	if len(req.Name) < 5 || len(req.Name) > 10 {
		return nil, fmt.Errorf("Name must be between 5 and 10 characters")
	}
	return h.domainAPIService.CreateProduct(ctx, req)
}

func (h *BffHandler) CreateOrder(ctx context.Context, req *bff_pb.NewOrder) (*bff_pb.OrderId, error) {
	if req.Count < 2 || req.Count > 100 {
		return nil, fmt.Errorf("Count must be between 2 and 100")
	}
	return h.domainAPIService.CreateOrder(ctx, req)
}
