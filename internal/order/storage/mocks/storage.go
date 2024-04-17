// Code generated by MockGen. DO NOT EDIT.
// Source: ./storage.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	model "homework/internal/order/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOrderStorage is a mock of OrderStorage interface.
type MockOrderStorage struct {
	ctrl     *gomock.Controller
	recorder *MockOrderStorageMockRecorder
}

// MockOrderStorageMockRecorder is the mock recorder for MockOrderStorage.
type MockOrderStorageMockRecorder struct {
	mock *MockOrderStorage
}

// NewMockOrderStorage creates a new mock instance.
func NewMockOrderStorage(ctrl *gomock.Controller) *MockOrderStorage {
	mock := &MockOrderStorage{ctrl: ctrl}
	mock.recorder = &MockOrderStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderStorage) EXPECT() *MockOrderStorageMockRecorder {
	return m.recorder
}

// AddOrder mocks base method.
func (m *MockOrderStorage) AddOrder(ctx context.Context, newOrder model.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrder", ctx, newOrder)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddOrder indicates an expected call of AddOrder.
func (mr *MockOrderStorageMockRecorder) AddOrder(ctx, newOrder interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrder", reflect.TypeOf((*MockOrderStorage)(nil).AddOrder), ctx, newOrder)
}

// DeleteOrder mocks base method.
func (m *MockOrderStorage) DeleteOrder(ctx context.Context, orderID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrder", ctx, orderID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrder indicates an expected call of DeleteOrder.
func (mr *MockOrderStorageMockRecorder) DeleteOrder(ctx, orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrder", reflect.TypeOf((*MockOrderStorage)(nil).DeleteOrder), ctx, orderID)
}

// DeleteOrdersByPPID mocks base method.
func (m *MockOrderStorage) DeleteOrdersByPPID(ctx context.Context, ppID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrdersByPPID", ctx, ppID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrdersByPPID indicates an expected call of DeleteOrdersByPPID.
func (mr *MockOrderStorageMockRecorder) DeleteOrdersByPPID(ctx, ppID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrdersByPPID", reflect.TypeOf((*MockOrderStorage)(nil).DeleteOrdersByPPID), ctx, ppID)
}

// GetOrderByID mocks base method.
func (m *MockOrderStorage) GetOrderByID(ctx context.Context, ID uint64) (*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByID", ctx, ID)
	ret0, _ := ret[0].(*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderByID indicates an expected call of GetOrderByID.
func (mr *MockOrderStorageMockRecorder) GetOrderByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByID", reflect.TypeOf((*MockOrderStorage)(nil).GetOrderByID), ctx, ID)
}

// GetOrderReturns mocks base method.
func (m *MockOrderStorage) GetOrderReturns(ctx context.Context) ([]model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderReturns", ctx)
	ret0, _ := ret[0].([]model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderReturns indicates an expected call of GetOrderReturns.
func (mr *MockOrderStorageMockRecorder) GetOrderReturns(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderReturns", reflect.TypeOf((*MockOrderStorage)(nil).GetOrderReturns), ctx)
}

// GetOrders mocks base method.
func (m *MockOrderStorage) GetOrders(ctx context.Context) ([]model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", ctx)
	ret0, _ := ret[0].([]model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockOrderStorageMockRecorder) GetOrders(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockOrderStorage)(nil).GetOrders), ctx)
}

// GetUserOrders mocks base method.
func (m *MockOrderStorage) GetUserOrders(ctx context.Context, clientID uint64) ([]model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserOrders", ctx, clientID)
	ret0, _ := ret[0].([]model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserOrders indicates an expected call of GetUserOrders.
func (mr *MockOrderStorageMockRecorder) GetUserOrders(ctx, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserOrders", reflect.TypeOf((*MockOrderStorage)(nil).GetUserOrders), ctx, clientID)
}

// IssueOrders mocks base method.
func (m *MockOrderStorage) IssueOrders(ctx context.Context, orderIDs map[uint64]struct{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IssueOrders", ctx, orderIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// IssueOrders indicates an expected call of IssueOrders.
func (mr *MockOrderStorageMockRecorder) IssueOrders(ctx, orderIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IssueOrders", reflect.TypeOf((*MockOrderStorage)(nil).IssueOrders), ctx, orderIDs)
}

// ReturnOrder mocks base method.
func (m *MockOrderStorage) ReturnOrder(ctx context.Context, clientID, orderID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReturnOrder", ctx, clientID, orderID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReturnOrder indicates an expected call of ReturnOrder.
func (mr *MockOrderStorageMockRecorder) ReturnOrder(ctx, clientID, orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReturnOrder", reflect.TypeOf((*MockOrderStorage)(nil).ReturnOrder), ctx, clientID, orderID)
}
