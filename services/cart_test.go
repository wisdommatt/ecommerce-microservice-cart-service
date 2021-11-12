package services

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/wisdommatt/ecommerce-microservice-cart-service/internal/cart"
	"github.com/wisdommatt/ecommerce-microservice-cart-service/mocks"
)

func TestCartServiceImpl_SaveCartItem(t *testing.T) {
	cartRepo := &mocks.Repository{}
	cartRepo.On("SaveCartItem", mock.Anything, &cart.CartItem{
		ProductSku: "invalid.sku",
		UserId:     "hello.world",
	}).Return(errors.New("an error occured"))

	cartRepo.On("SaveCartItem", mock.Anything, &cart.CartItem{
		ProductSku: "valid.sku",
		UserId:     "hello.world",
	}).Return(nil).Run(func(args mock.Arguments) {
		item := args[1].(*cart.CartItem)
		item.Quantity = 30
	})

	type args struct {
		item *cart.CartItem
	}
	tests := []struct {
		name    string
		args    args
		want    *cart.CartItem
		wantErr bool
	}{
		{
			name: "SaveCartItem repo implementation with error",
			args: args{item: &cart.CartItem{
				ProductSku: "invalid.sku",
				UserId:     "hello.world",
			}},
			wantErr: true,
		},
		{
			name: "SaveCartItem repo implementation without error",
			args: args{item: &cart.CartItem{
				ProductSku: "valid.sku",
				UserId:     "hello.world",
			}},
			want: &cart.CartItem{
				ProductSku: "valid.sku",
				UserId:     "hello.world",
				Quantity:   30,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewCartService(cartRepo)
			got, err := s.SaveCartItem(context.Background(), tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("CartServiceImpl.SaveCartItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartServiceImpl.SaveCartItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
