package mocktest

import (
	pbp "EXAM3/api-gateway/genproto/product_service"
	"context"

	"google.golang.org/grpc"
)

type ProductServiceClient interface {
	CreateProduct(ctx context.Context, in *pbp.Product, opts ...grpc.CallOption) (*pbp.Product, error)
	GetProductById(ctx context.Context, in *pbp.ProductId, opts ...grpc.CallOption) (*pbp.Product, error)
	UpdateProduct(ctx context.Context, in *pbp.Product, opts ...grpc.CallOption) (*pbp.Product, error)
	DeleteProduct(ctx context.Context, in *pbp.ProductId, opts ...grpc.CallOption) (*pbp.Status, error)
	ListProducts(ctx context.Context, in *pbp.GetAllProductRequest, opts ...grpc.CallOption) (*pbp.GetAllProductResponse, error)
}

type productServiceClient struct {
}

func NewProductServiceClient() ProductServiceClient {
	return &productServiceClient{}
}

func (c *productServiceClient) CreateProduct(ctx context.Context, in *pbp.Product, opts ...grpc.CallOption) (*pbp.Product, error) {
	return in, nil
}

func (c *productServiceClient) GetProductById(ctx context.Context, in *pbp.ProductId, opts ...grpc.CallOption) (*pbp.Product, error) {

	return &pbp.Product{
		Id:          "yewruoe",
		Name:        "Product name",
		Description: "Product description",
		Price:       10.1,
		Amount:      5,
	}, nil
}

func (c *productServiceClient) UpdateProduct(ctx context.Context, in *pbp.Product, opts ...grpc.CallOption) (*pbp.Product, error) {
	return &pbp.Product{
		Id:          "yewruoe",
		Name:        "Product name",
		Description: "Product description",
		Price:       10.1,
		Amount:      5,
	}, nil
}

func (c *productServiceClient) DeleteProduct(ctx context.Context, in *pbp.ProductId, opts ...grpc.CallOption) (*pbp.Status, error) {
	return &pbp.Status{
		Success: true,
	}, nil
}

func (c *productServiceClient) ListProducts(ctx context.Context, in *pbp.GetAllProductRequest, opts ...grpc.CallOption) (*pbp.GetAllProductResponse, error) {
	pr := pbp.Product{
		Id:          "yewruoe",
		Name:        "Product name",
		Description: "Product description",
		Price:       10.1,
		Amount:      5,
	}
	return &pbp.GetAllProductResponse{
		Count: 3,
		Products: []*pbp.Product{
			&pr,
			&pr,
			&pr,
		},
	}, nil
}
