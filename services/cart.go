package services

import (
	"context"
	"errors"

	"github.com/wisdommatt/ecommerce-microservice-cart-service/internal/cart"
)

// CartService is the interface that describes a cart service
// object.
type CartService interface {
	SaveCartItem(ctx context.Context, item *cart.CartItem) (*cart.CartItem, error)
}

type CartServiceImpl struct {
	cartRepo cart.Repository
}

// NewCartService returns a new cart service object.
func NewCartService(cartRepo cart.Repository) *CartServiceImpl {
	return &CartServiceImpl{
		cartRepo: cartRepo,
	}
}

// SaveCartItem is the service handler to save an item to cart.
func (s *CartServiceImpl) SaveCartItem(ctx context.Context, item *cart.CartItem) (*cart.CartItem, error) {
	err := s.cartRepo.SaveCartItem(ctx, item)
	if err != nil {
		return nil, errors.New("an error occured, please try again later")
	}
	return item, nil
}
