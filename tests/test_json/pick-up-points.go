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
)
