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
	BulkRemoveItemsFromUserCart(ctx context.Context, userId string, itemIds []string) ([]cart.CartItem, error)
}

type CartServiceImpl struct {
	cartRepo cart.Repository
	tracer   opentracing.Tracer
}

// NewCartService returns a new cart service object.
func NewCartService(cartRepo cart.Repository, tracer opentracing.Tracer) *CartServiceImpl {
	return &CartServiceImpl{
		cartRepo: cartRepo,
		tracer:   tracer,
	}
}

// SaveCartItem is the service handler to save an item to cart.
func (s *CartServiceImpl) SaveCartItem(ctx context.Context, item *cart.CartItem) (*cart.CartItem, error) {
	span := s.tracer.StartSpan("SaveCartItem")
	if opentracing.SpanFromContext(ctx) != nil {
		span = s.tracer.StartSpan(
			"SaveCartItem",
			opentracing.ChildOf(opentracing.SpanFromContext(ctx).Context()),
		)
	}
	defer span.Finish()

	err := s.cartRepo.SaveCartItem(opentracing.ContextWithSpan(ctx, span), item)
	if err != nil {
		return nil, errors.New("an error occured, please try again later")
	}
	return item, nil
}

func (s *CartServiceImpl) GetUserCartItems(ctx context.Context, userId string) ([]cart.CartItem, error) {
	span := s.tracer.StartSpan("GetUserCartItems")
	if opentracing.SpanFromContext(ctx) != nil {
		span = s.tracer.StartSpan(
			"GetUserCartItems",
			opentracing.ChildOf(opentracing.SpanFromContext(ctx).Context()),
		)
	}
	defer span.Finish()

	if userId == "" {
		ext.Error.Set(span, true)
		span.LogFields(log.Error(errors.New("userId should not be empty")))
		return nil, errors.New("userId must be provided")
	}
	items, err := s.cartRepo.GetUserCartItems(opentracing.ContextWithSpan(ctx, span), userId)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *CartServiceImpl) BulkRemoveItemsFromUserCart(
	ctx context.Context, userId string, itemIds []string,
) ([]cart.CartItem, error) {

	span := s.tracer.StartSpan("BulkRemoveItemsFromUserCart")
	if opentracing.SpanFromContext(ctx) != nil {
		span = s.tracer.StartSpan(
			"BulkRemoveItemsFromUserCart",
			opentracing.ChildOf(opentracing.SpanFromContext(ctx).Context()),
		)
	}
	defer span.Finish()

	if userId == "" || len(itemIds) == 0 {
		ext.Error.Set(span, true)
		span.LogFields(log.Error(errors.New("all fields are required")))
		return nil, errors.New("all fields are required")
	}
	err := s.cartRepo.BulkRemoveItemsFromUserCart(opentracing.ContextWithSpan(ctx, span), userId, itemIds)
	if err != nil {
		return nil, errors.New("an error occured, please try again later")
	}
	return s.GetUserCartItems(opentracing.ContextWithSpan(ctx, span), userId)
}
