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

	InternalError = `{"result":"internal error"}`

	InvalidInput = `{"result":"invalid json passed"}`

	SuccessResult = `{"result":"success"}`

	ValidPPResponse = `{"ID":5000,"Name":"PickUpPoint1","Address":{"Region":"Курская область","City":"Курск","Street":"Студенческая","HouseNum":"2A"},"PhoneNumber":"88005553535"}`
)
