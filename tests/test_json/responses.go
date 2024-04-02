package test_json

const (
	InValidPPRequest = `{
		"name": "PickUpPoint1",
			"address": {
			"region": "Курская область",
				"city": "Курск",
				"street": "Студенческая"
		},
		"phone_number": "88005553535"}`

	ValidPPAddRequest = `{
		"name": "PickUpPoint1",
			"address": {
			"region": "Курская область",
  			"city": "Курск",
			"street": "Студенческая",
			"house_num": "2A"
		},
		"phone_number": "88005553535"}`

	ValidPPAddRequestUnique = `{
		"name": "PickUpPointNew",
			"address": {
			"region": "Курская область",
  			"city": "Курск",
			"street": "Студенческая",
			"house_num": "2A"
		},
		"phone_number": "88005553535"}`

	ValidPPUpdateRequest = `{
		"id": 5000,
		"name": "PickUpPoint1",
			"address": {
			"region": "Курская область",
  			"city": "Курск",
			"street": "Студенческая",
			"house_num": "2A"
		},
		"phone_number": "88005553535"}`

	ValidPPUpdateRequestNotExists = `{
		"id": 5020,
		"name": "PickUpPoint1",
			"address": {
			"region": "Курская область",
  			"city": "Курск",
			"street": "Студенческая",
			"house_num": "2A"
		},
		"phone_number": "88005553535"}`

	ValidPPUpdateRequestNameAlreadyExists = `{
		"id": 5000,
		"name": "PickUpPoint2",
			"address": {
			"region": "Курская область",
  			"city": "Курск",
			"street": "Студенческая",
			"house_num": "2A"
		},
		"phone_number": "88005553535"}`

	InternalError = `{"result":"internal error"}`

	InvalidInput = `{"result":"invalid json passed"}`

	SuccessResult = `{"result":"success"}`

	ValidPPResponse = `{"ID":5000,"Name":"PickUpPoint1","Address":{"Region":"Курская область","City":"Курск","Street":"Студенческая","HouseNum":"2A"},"PhoneNumber":"88005553535"}`

	ValidPPResponseAdd = `{"ID":5010,"Name":"PickUpPointNew","Address":{"Region":"Курская область","City":"Курск","Street":"Студенческая","HouseNum":"2A"},"PhoneNumber":"88005553535"}`

	ValidPPSliceResponse = `[{"ID":5000,"Name":"PickUpPoint1","Address":{"Region":"Курская область","City":"Курск","Street":"Студенческая","HouseNum":"2A"},"PhoneNumber":"88005553535"},{"ID":5001,"Name":"PickUpPoint2","Address":{"Region":"Курская область","City":"Курск","Street":"Студенческая","HouseNum":"2A"},"PhoneNumber":"88005553535"}]`
)
