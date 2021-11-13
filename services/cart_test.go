package services

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/opentracing/opentracing-go"
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
			s := NewCartService(cartRepo, &opentracing.NoopTracer{})
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
			s := NewCartService(cartRepo, &opentracing.NoopTracer{})
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

func TestCartServiceImpl_BulkRemoveItemsFromUserCart(t *testing.T) {
	cartRepo := &mocks.Repository{}
	cartRepo.On("BulkRemoveItemsFromUserCart", mock.Anything, "invalid.user", []string{"123", "456"}).
		Return(errors.New("an error occured"))

	cartRepo.On("BulkRemoveItemsFromUserCart", mock.Anything, mock.Anything, []string{"123", "456"}).
		Return(nil)

	cartRepo.On("GetUserCartItems", mock.Anything, "err.user").
		Return(nil, errors.New("an error occured"))

	cartRepo.On("GetUserCartItems", mock.Anything, "noerr.user").
		Return([]cart.CartItem{
			{ProductSku: "sku.1"}, {ProductSku: "sku.2"}, {ProductSku: "sku.3"},
		}, nil)

	type args struct {
		userId  string
		itemIds []string
	}
	tests := []struct {
		name    string
		args    args
		want    []cart.CartItem
		wantErr bool
	}{
		{
			name:    "empty userId",
			args:    args{userId: "", itemIds: []string{"hello"}},
			wantErr: true,
		},
		{
			name:    "empty itemIds",
			args:    args{userId: "hello", itemIds: []string{}},
			wantErr: true,
		},
		{
			name:    "empty userId & itemIds",
			args:    args{userId: "", itemIds: []string{}},
			wantErr: true,
		},
		{
			name:    "BulkRemoveItemsFromUserCart repo implementation with error",
			args:    args{userId: "invalid.user", itemIds: []string{"123", "456"}},
			wantErr: true,
		},
		{
			name:    "GetUserCartItems repo implementation with error",
			args:    args{userId: "err.user", itemIds: []string{"123", "456"}},
			wantErr: true,
		},
		{
			name: "GetUserCartItems repo implementation without error",
			args: args{userId: "noerr.user", itemIds: []string{"123", "456"}},
			want: []cart.CartItem{
				{ProductSku: "sku.1"}, {ProductSku: "sku.2"}, {ProductSku: "sku.3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewCartService(cartRepo, &opentracing.NoopTracer{})
			got, err := s.BulkRemoveItemsFromUserCart(context.Background(), tt.args.userId, tt.args.itemIds)
			if (err != nil) != tt.wantErr {
				t.Errorf("CartServiceImpl.BulkRemoveItemsFromUserCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartServiceImpl.BulkRemoveItemsFromUserCart() = %v, want %v", got, tt.want)
			}
		})
	}
}
