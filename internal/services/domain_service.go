package services

import (
	"context"
	pb "specmatic-order-bff-grpc-go/pkg/api/proto_files"

	"google.golang.org/grpc"
)

type DomainAPIService struct {
	orderServiceClient   pb.OrderServiceClient
	productServiceClient pb.ProductServiceClient
}

func NewDomainAPIService(orderConn, productConn *grpc.ClientConn) *DomainAPIService {
	return &DomainAPIService{
		orderServiceClient:   pb.NewOrderServiceClient(orderConn),
		productServiceClient: pb.NewProductServiceClient(productConn),
	}
}

func (s *DomainAPIService) CreateOrder(ctx context.Context, newOrder *pb.NewOrder) (*pb.OrderId, error) {
	domainOrder := &pb.NewOrder{
		ProductId: newOrder.ProductId,
		Count:     newOrder.Count,
		Status:    pb.OrderStatus_PENDING,
	}
	orderId, err := s.orderServiceClient.AddOrder(ctx, domainOrder)
	if err != nil {
		return nil, err
	}
	return &pb.OrderId{Id: orderId.Id}, nil
}

func (s *DomainAPIService) FindProducts(ctx context.Context, req *pb.FindAvailableProductsRequest) (*pb.ProductListResponse, error) {
	domainReq := &pb.ProductSearchRequest{
		Type: pb.ProductType(req.Type),
	}
	domainResp, err := s.productServiceClient.SearchProducts(ctx, domainReq)
	if err != nil {
		return nil, err
	}

	products := make([]*pb.Product, len(domainResp.Products))
	for i, p := range domainResp.Products {
		products[i] = &pb.Product{
			Id:        p.Id,
			Name:      p.Name,
			Type:      pb.ProductType(p.Type),
			Inventory: p.Inventory,
		}
	}
	return &pb.ProductListResponse{Products: products}, nil
}

func (s *DomainAPIService) CreateProduct(ctx context.Context, newProduct *pb.NewProduct) (*pb.ProductId, error) {
	domainProduct := &pb.NewProduct{
		Name:      newProduct.Name,
		Type:      pb.ProductType(newProduct.Type),
		Inventory: newProduct.Inventory,
	}
	productId, err := s.productServiceClient.AddProduct(ctx, domainProduct)
	if err != nil {
		return nil, err
	}
	return &pb.ProductId{Id: productId.Id}, nil
}
