package handlers

import (
	"context"
	"log"
	"specmatic-order-bff-grpc-go/internal/services"
	bff_pb "specmatic-order-bff-grpc-go/pkg/api/in/specmatic/examples/store/order_bff_grpc"

	"specmatic-order-bff-grpc-go/internal/utils"
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
	if err := utils.ValidateReq(req); err != nil {
		log.Printf("FindAvailableProducts validation error: %v \n \n", err)
		return nil, err
	}
	return h.domainAPIService.FindProducts(ctx, req)
}

func (h *BffHandler) CreateProduct(ctx context.Context, req *bff_pb.NewProduct) (*bff_pb.ProductId, error) {
	if err := utils.ValidateReq(req); err != nil {
		log.Printf("Create Product validation error: %v \n \n", err)
		return nil, err
	}
	return h.domainAPIService.CreateProduct(ctx, req)
}

func (h *BffHandler) CreateOrder(ctx context.Context, req *bff_pb.NewOrder) (*bff_pb.OrderId, error) {
	if err := utils.ValidateReq(req); err != nil {
		log.Printf("Create Order validation error: %v \n \n", err)
		return nil, err
	}
	return h.domainAPIService.CreateOrder(ctx, req)
}
