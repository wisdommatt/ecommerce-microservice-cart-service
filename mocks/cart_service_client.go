// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	proto "github.com/wisdommatt/ecommerce-microservice-cart-service/grpc/proto"
)

// CartServiceClient is an autogenerated mock type for the CartServiceClient type
type CartServiceClient struct {
	mock.Mock
}

// AddToCart provides a mock function with given fields: ctx, in, opts
func (_m *CartServiceClient) AddToCart(ctx context.Context, in *proto.NewCartItem, opts ...grpc.CallOption) (*proto.CartItem, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.CartItem
	if rf, ok := ret.Get(0).(func(context.Context, *proto.NewCartItem, ...grpc.CallOption) *proto.CartItem); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.CartItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.NewCartItem, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserCart provides a mock function with given fields: ctx, in, opts
func (_m *CartServiceClient) GetUserCart(ctx context.Context, in *proto.GetUserCartInput, opts ...grpc.CallOption) (*proto.GetUserCartResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GetUserCartResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.GetUserCartInput, ...grpc.CallOption) *proto.GetUserCartResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GetUserCartResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.GetUserCartInput, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveItemsFromCart provides a mock function with given fields: ctx, in, opts
func (_m *CartServiceClient) RemoveItemsFromCart(ctx context.Context, in *proto.RemoveItemsFromCartInput, opts ...grpc.CallOption) (*proto.RemoveItemsFromCartResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.RemoveItemsFromCartResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.RemoveItemsFromCartInput, ...grpc.CallOption) *proto.RemoveItemsFromCartResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.RemoveItemsFromCartResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.RemoveItemsFromCartInput, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
