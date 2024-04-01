package storage

import (
	"sync"

	"homework/internal/pick-up_point/model"
	"homework/tests/fixtures"
	"homework/tests/states"
)

type pickUpPointStorageFixtures struct {
	st FilePPStorage
}

func setUp() pickUpPointStorageFixtures {
	st := FilePPStorage{
		cache:  []model.PickUpPoint{fixtures.PickUpPoint().Valid().V()},
		mu:     &sync.RWMutex{},
		nextID: states.PPID,
	}
	return pickUpPointStorageFixtures{
		st: st,
	}
}
