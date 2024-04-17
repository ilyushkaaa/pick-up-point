package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	mockStorageOrder "homework/internal/order/storage/mocks"
	mockStoragePP "homework/internal/pick-up_point/storage/mocks"
)

type pickUpPointServiceFixtures struct {
	ctrl        *gomock.Controller
	srv         PPService
	mockStorage *mockStoragePP.MockPPStorage
}

func setUp(t *testing.T) pickUpPointServiceFixtures {
	ctrl := gomock.NewController(t)
	mockPPStorage := mockStoragePP.NewMockPPStorage(ctrl)
	mockOrderStorage := mockStorageOrder.NewMockOrderStorage(ctrl)
	srv := PPService{
		orderStorage: mockOrderStorage,
		ppStorage:    mockPPStorage,
	}
	return pickUpPointServiceFixtures{
		ctrl:        ctrl,
		srv:         srv,
		mockStorage: mockPPStorage,
	}
}
