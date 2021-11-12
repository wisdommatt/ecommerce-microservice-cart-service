package servers

import (
	"strconv"

	"github.com/wisdommatt/ecommerce-microservice-cart-service/grpc/proto"
	"github.com/wisdommatt/ecommerce-microservice-cart-service/internal/cart"
)

func ProtoNewCartItemToInternal(item *proto.NewCartItem) *cart.CartItem {
	return &cart.CartItem{
		ProductSku: item.ProductSku,
		UserId:     item.UserId,
		Quantity:   int(item.Quantity),
	}
}

func InternalCartItemToProto(item *cart.CartItem) *proto.CartItem {
	return &proto.CartItem{
		Id:         strconv.Itoa(item.ID),
		ProductSku: item.ProductSku,
		UserId:     item.UserId,
		Quantity:   int32(item.Quantity),
	}
}
