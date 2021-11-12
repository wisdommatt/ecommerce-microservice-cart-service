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

func TestCartServiceImpl_GetUserCartItems(t *testing.T) {
	cartRepo := &mocks.Repository{}
	cartRepo.On("GetUserCartItems", mock.Anything, "invalid.user").Return(nil, errors.New("an error occured"))
	cartRepo.On("GetUserCartItems", mock.Anything, "valid.user").Return([]cart.CartItem{
		{ProductSku: "sku.1"}, {ProductSku: "sku.2"}, {ProductSku: "sku.3"},
	}, nil)

	type args struct {
		userId string
	}
	tests := []struct {
		name    string
		args    args
		want    []cart.CartItem
		wantErr bool
	}{
		{
			name:    "empty userId",
			wantErr: true,
		},
		{
			name:    "GetUserCartItems repo implementation with error",
			args:    args{userId: "invalid.user"},
			wantErr: true,
		},
		{
			name: "GetUserCartItems repo implementation without error",
			args: args{userId: "valid.user"},
			want: []cart.CartItem{
				{ProductSku: "sku.1"}, {ProductSku: "sku.2"}, {ProductSku: "sku.3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewCartService(cartRepo)
			got, err := s.GetUserCartItems(context.Background(), tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("CartServiceImpl.GetUserCartItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartServiceImpl.GetUserCartItems() = %v, want %v", got, tt.want)
			}
		})
	}
}
