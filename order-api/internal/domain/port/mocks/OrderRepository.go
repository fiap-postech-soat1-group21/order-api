// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/fiap-postech-soat1-group21/order-api/order-api/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// OrderRepository is an autogenerated mock type for the OrderRepository type
type OrderRepository struct {
	mock.Mock
}

// CreateOrder provides a mock function with given fields: _a0, _a1
func (_m *OrderRepository) CreateOrder(_a0 context.Context, _a1 *entity.Order) (*entity.Order, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrder")
	}

	var r0 *entity.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Order) (*entity.Order, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Order) *entity.Order); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.Order) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrderItems provides a mock function with given fields: ctx, order
func (_m *OrderRepository) CreateOrderItems(ctx context.Context, order []*entity.OrderItems) ([]*entity.OrderItems, error) {
	ret := _m.Called(ctx, order)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrderItems")
	}

	var r0 []*entity.OrderItems
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*entity.OrderItems) ([]*entity.OrderItems, error)); ok {
		return rf(ctx, order)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*entity.OrderItems) []*entity.OrderItems); ok {
		r0 = rf(ctx, order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.OrderItems)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*entity.OrderItems) error); ok {
		r1 = rf(ctx, order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderItems provides a mock function with given fields: ctx, orderID
func (_m *OrderRepository) GetOrderItems(ctx context.Context, orderID uuid.UUID) ([]*entity.OrderItems, error) {
	ret := _m.Called(ctx, orderID)

	if len(ret) == 0 {
		panic("no return value specified for GetOrderItems")
	}

	var r0 []*entity.OrderItems
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]*entity.OrderItems, error)); ok {
		return rf(ctx, orderID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []*entity.OrderItems); ok {
		r0 = rf(ctx, orderID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.OrderItems)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0
func (_m *OrderRepository) List(_a0 context.Context) (*entity.OrderResponseList, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *entity.OrderResponseList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*entity.OrderResponseList, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *entity.OrderResponseList); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.OrderResponseList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOrderStatus provides a mock function with given fields: ctx, orderID, status
func (_m *OrderRepository) UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status string) error {
	ret := _m.Called(ctx, orderID, status)

	if len(ret) == 0 {
		panic("no return value specified for UpdateOrderStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) error); ok {
		r0 = rf(ctx, orderID, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewOrderRepository creates a new instance of OrderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrderRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrderRepository {
	mock := &OrderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
