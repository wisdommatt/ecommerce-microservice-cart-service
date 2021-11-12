package services

import (
	"context"
	"errors"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"github.com/wisdommatt/ecommerce-microservice-cart-service/internal/cart"
)

// CartService is the interface that describes a cart service
// object.
type CartService interface {
	SaveCartItem(ctx context.Context, item *cart.CartItem) (*cart.CartItem, error)
	GetUserCartItems(ctx context.Context, userId string) ([]cart.CartItem, error)
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

func (s *CartServiceImpl) GetUserCartItems(ctx context.Context, userId string) ([]cart.CartItem, error) {
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		span = opentracing.StartSpan("services.GetUserCartItems")
	}
	if userId == "" {
		ext.Error.Set(span, true)
		span.LogFields(log.Error(errors.New("userId should not be empty")))
		return nil, errors.New("userId must be provided")
	}
	items, err := s.cartRepo.GetUserCartItems(ctx, userId)
	if err != nil {
		return nil, err
	}
	return items, nil
}
