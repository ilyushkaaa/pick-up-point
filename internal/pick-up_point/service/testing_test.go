package service

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"go.uber.org/zap"
	cacheInMemory "homework/internal/cache/in_memory"
	mockStorageOrder "homework/internal/order/storage/mocks"
	mockStoragePP "homework/internal/pick-up_point/storage/mocks"
)

type pickUpPointServiceFixtures struct {
	ctrl        *gomock.Controller
	srv         *PPService
	mockStorage *mockStoragePP.MockPPStorage
}

func setUp(t *testing.T) pickUpPointServiceFixtures {
	ctrl := gomock.NewController(t)
	mockPPStorage := mockStoragePP.NewMockPPStorage(ctrl)
	mockOrderStorage := mockStorageOrder.NewMockOrderStorage(ctrl)
	tm := &fakeTransactionManager{}
	logger := zap.NewNop().Sugar()
	imMemoryCache := cacheInMemory.New(logger, time.Minute)
	srv := New(mockPPStorage, mockOrderStorage, tm, imMemoryCache)

	return pickUpPointServiceFixtures{
		ctrl:        ctrl,
		srv:         srv,
		mockStorage: mockPPStorage,
	}
}

type fakeTransactionManager struct{}

func (t *fakeTransactionManager) RunSerializable(ctx context.Context, f func(ctxTX context.Context) error) error {
	return f(ctx)
}
func (t *fakeTransactionManager) RunReadCommitted(ctx context.Context, f func(ctxTX context.Context) error) error {
	return f(ctx)
}
func (t *fakeTransactionManager) RunRepeatableRead(ctx context.Context, f func(ctxTX context.Context) error) error {
	return f(ctx)
}
func (t *fakeTransactionManager) RunReadUnCommitted(ctx context.Context, f func(ctxTX context.Context) error) error {
	return f(ctx)
}
