package service

import (
	"context"

	cartpb "github.com/e-commerce-microservice/shopping-cart/gen"
	"github.com/e-commerce-microservice/shopping-cart/internals/repo"
)

type CartService struct {
	repo repo.ICartRepo
	cartpb.UnimplementedShoppingCartServiceServer
}

func NewCartService(repo repo.ICartRepo) *CartService {
	return &CartService{repo: repo}
}

func (c *CartService) CreateShoppingCart(ctx context.Context, req *cartpb.CreateShoppingCartRequest) (*cartpb.CreateShoppingCartResponse, error) {

	id, err := c.repo.CreateCart(ctx, uint(req.UserId))

	if err != nil {
		return &cartpb.CreateShoppingCartResponse{Success: false}, err
	}

	return &cartpb.CreateShoppingCartResponse{Success: true, CartId: uint32(id)}, nil
}

func (c *CartService) AddItemToCart(ctx context.Context, req *cartpb.AddItemToCartRequest) (*cartpb.AddItemToCartResponse, error) {
	err := c.repo.AddItemToCart(ctx, uint(req.CartId), uint(req.ProductId), int(req.Quantity))
	if err != nil {
		return &cartpb.AddItemToCartResponse{Success: false}, err
	}
	return &cartpb.AddItemToCartResponse{Success: true}, nil
}
