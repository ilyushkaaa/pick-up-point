package fixtures

import (
	"homework/internal/pick-up_point/model"
	"homework/tests/states"
)

type PickUpPointBuilder struct {
	instance *model.PickUpPoint
}

func PickUpPoint() *PickUpPointBuilder {
	return &PickUpPointBuilder{instance: &model.PickUpPoint{}}
}

func (b *PickUpPointBuilder) ID(v uint64) *PickUpPointBuilder {
	b.instance.ID = v
	return b
}

func (b *PickUpPointBuilder) Name(v string) *PickUpPointBuilder {
	b.instance.Name = v
	return b
}

func (b *PickUpPointBuilder) Region(v string) *PickUpPointBuilder {
	b.instance.Address.Region = v
	return b
}

func (b *PickUpPointBuilder) City(v string) *PickUpPointBuilder {
	b.instance.Address.City = v
	return b
}

func (b *PickUpPointBuilder) Street(v string) *PickUpPointBuilder {
	b.instance.Address.Street = v
	return b
}

func (b *PickUpPointBuilder) HouseNum(v string) *PickUpPointBuilder {
	b.instance.Address.HouseNum = v
	return b
}
func (b *PickUpPointBuilder) PhoneNumber(v string) *PickUpPointBuilder {
	b.instance.PhoneNumber = v
	return b
}

func (b *PickUpPointBuilder) P() *model.PickUpPoint {
	return b.instance
}

func (b *PickUpPointBuilder) V() model.PickUpPoint {
	return *b.instance
}

func (b *PickUpPointBuilder) ValidWithoutID() *PickUpPointBuilder {
	return PickUpPoint().Name(states.PPName).Region(states.PPRegion).City(states.PPCity).Street(states.PPStreet).
		HouseNum(states.PPHouseNum).PhoneNumber(states.PPPhoneNumber)
}

func (b *PickUpPointBuilder) Valid() *PickUpPointBuilder {
	return b.ValidWithoutID().ID(states.PPID)
}
