package api

import (
	"testing"
)

type TestJsonStruct struct {
	Email string `json:"name"`
	Age   int    `json:"age"`
	Phone int64  `json:"phone"`
}

func TestApiResponse(t *testing.T) {

	apiResponse := &ApiResponse{Message: "Something Up M8",
		Payload: &TestJsonStruct{
			Email: "test@test.com",
			Age:   100,
			Phone: 125121251,
		}}

	if apiResponse.statusCode != 0 {
		t.Error("Expected statusCode to be 0, but got", apiResponse.statusCode, "instead")
	}

	apiResponse.Status(401)

	if apiResponse.statusCode != 401 {
		t.Error("Expected statusCode to be 401, but got", apiResponse.statusCode, "instead")
	}

	if apiResponse.Message != "Something Up M8" {
		t.Error("Expected Message to be 'Something Up M8' but got", apiResponse.Message, "instead")
	}

	testJsonValue := apiResponse.Payload.(*TestJsonStruct)

	if testJsonValue.Age != 100 {
		t.Error("Expected Payload's Age value to be 100 but got", testJsonValue.Age, "instead")
	}

	if testJsonValue.Email != "test@test.com" {
		t.Error("Expected Payload's Email value to be test@test.com but got", testJsonValue.Email, "instead")
	}

	if testJsonValue.Phone != 125121251 {
		t.Error("Expected Payload's Phone value to be 125121251 but got", testJsonValue.Phone, "instead")
	}
}
