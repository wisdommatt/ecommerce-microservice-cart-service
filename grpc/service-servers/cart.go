package servers

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/wisdommatt/ecommerce-microservice-cart-service/grpc/proto"
	"github.com/wisdommatt/ecommerce-microservice-cart-service/services"
)

type CartServer struct {
	proto.UnimplementedCartServiceServer
	cartService services.CartService
}

// NewCartServer returns a new cart grpc service server object.
func NewCartServer(cartService services.CartService) *CartServer {
	return &CartServer{
		cartService: cartService,
	}
}

// AddToCart is the grpc handler to add item to cart.
func (s *CartServer) AddToCart(ctx context.Context, req *proto.NewCartItem) (*proto.CartItem, error) {
	span := opentracing.StartSpan("AddToCart")
	defer span.Finish()
	ext.SpanKindRPCServer.Set(span)
	span.SetTag("param.req", req)

	ctx = opentracing.ContextWithSpan(ctx, span)
	newCartItem, err := s.cartService.SaveCartItem(ctx, ProtoNewCartItemToInternal(req))
	if err != nil {
		return nil, err
	}
	return InternalCartItemToProto(newCartItem), nil
}