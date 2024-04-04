package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_storage "homework/internal/pick-up_point/storage/mocks"
)

type pickUpPointServiceFixtures struct {
	ctrl        *gomock.Controller
	srv         PPService
	mockStorage *mock_storage.MockPPStorage
}

func setUp(t *testing.T) pickUpPointServiceFixtures {
	ctrl := gomock.NewController(t)
	mockPPStorage := mock_storage.NewMockPPStorage(ctrl)
	srv := PPService{mockPPStorage}
	return pickUpPointServiceFixtures{
		ctrl:        ctrl,
		srv:         srv,
		mockStorage: mockPPStorage,
	}
}

func (a *pickUpPointServiceFixtures) tearDown() {
	a.ctrl.Finish()
}
