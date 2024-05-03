package test_pb

import pb "homework/internal/pb/pick-up_point"

var ValidPPAddRequest = pb.PickUpPointAdd{
	Name: "PickUpPoint1",
	Address: &pb.AddressDTO{
		Region:   "Курская область",
		City:     "Курск",
		Street:   "Студенческая",
		HouseNum: "2A",
	},
	PhoneNumber: "88005553535",
}

var ValidPPAddRequestUnique = pb.PickUpPointAdd{
	Name: "PickUpPointNew",
	Address: &pb.AddressDTO{
		Region:   "Курская область",
		City:     "Курск",
		Street:   "Студенческая",
		HouseNum: "2A",
	},
	PhoneNumber: "88005553535",
}

var AddedPP = pb.PickUpPoint{
	Id:   1,
	Name: "PickUpPointNew",
	Address: &pb.Address{
		Region:   "Курская область",
		City:     "Курск",
		Street:   "Студенческая",
		HouseNum: "2A",
	},
	PhoneNumber: "88005553535",
}

var DeletePPRequestNotExist = pb.DeletePPRequest{
	Id: 10000,
}

var DeletePPRequestOk = pb.DeletePPRequest{
	Id: 5000,
}

var DeleteSuccessResult = pb.DeleteResponse{
	Result: "success",
}

var GetByIDRequestNotExist = pb.GetByIDRequest{
	Id: 10000,
}

var GetByIDRequestOk = pb.GetByIDRequest{
	Id: 5000,
}

var GetPPByIDSuccess = pb.PickUpPoint{
	Id:   5000,
	Name: "PickUpPoint1",
	Address: &pb.Address{
		Region:   "Курская область",
		City:     "Курск",
		Street:   "Студенческая",
		HouseNum: "2A",
	},
	PhoneNumber: "88005553535",
}

var GetAllPP = []*pb.PickUpPoint{
	{
		Id:   5000,
		Name: "PickUpPoint1",
		Address: &pb.Address{
			Region:   "Курская область",
			City:     "Курск",
			Street:   "Студенческая",
			HouseNum: "2A",
		},
		PhoneNumber: "88005553535",
	},
	{
		Id:   5001,
		Name: "PickUpPoint2",
		Address: &pb.Address{
			Region:   "Курская область",
			City:     "Курск",
			Street:   "Студенческая",
			HouseNum: "2A",
		},
		PhoneNumber: "88005553535",
	},
}

var UpdatePPNotExist = pb.PickUpPointUpdate{
	Id:   5020,
	Name: "PickUpPoint1",
	Address: &pb.AddressDTO{
		Region:   "Курская область",
		City:     "Курск",
		Street:   "Студенческая",
		HouseNum: "2A",
	},
	PhoneNumber: "88005553535",
}

var UpdatePPNameAlreadyExists = pb.PickUpPointUpdate{
	Id:   5000,
	Name: "PickUpPoint2",
	Address: &pb.AddressDTO{
		Region:   "Курская область",
		City:     "Курск",
		Street:   "Студенческая",
		HouseNum: "2A",
	},
	PhoneNumber: "88005553535",
}

var UpdatePPOk = pb.PickUpPointUpdate{
	Id:   5000,
	Name: "PickUpPoint10",
	Address: &pb.AddressDTO{
		Region:   "Курская область",
		City:     "Курск",
		Street:   "Студенческая",
		HouseNum: "2A",
	},
	PhoneNumber: "88005553535",
}
