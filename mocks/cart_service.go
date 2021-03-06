// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	cart "github.com/wisdommatt/ecommerce-microservice-cart-service/internal/cart"

	mock "github.com/stretchr/testify/mock"
)

// CartService is an autogenerated mock type for the CartService type
type CartService struct {
	mock.Mock
}

// GetUserCartItems provides a mock function with given fields: ctx, userId
func (_m *CartService) GetUserCartItems(ctx context.Context, userId string) ([]cart.CartItem, error) {
	ret := _m.Called(ctx, userId)

	var r0 []cart.CartItem
	if rf, ok := ret.Get(0).(func(context.Context, string) []cart.CartItem); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cart.CartItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveCartItem provides a mock function with given fields: ctx, item
func (_m *CartService) SaveCartItem(ctx context.Context, item *cart.CartItem) (*cart.CartItem, error) {
	ret := _m.Called(ctx, item)

	var r0 *cart.CartItem
	if rf, ok := ret.Get(0).(func(context.Context, *cart.CartItem) *cart.CartItem); ok {
		r0 = rf(ctx, item)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.CartItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cart.CartItem) error); ok {
		r1 = rf(ctx, item)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
