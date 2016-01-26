package goutyl

import (
	"testing"
)

func TestJson(t *testing.T) {
	jsonByteArr, err := Json(&TestJsonStruct{
		Email: "test2@test.com",
		Age:   42,
		Phone: 12512512,
	})

	if err != nil {
		t.Error(err.Error())
	}

	if len(jsonByteArr) == 0 {
		t.Error("Expected jsonByteArr to not be empty")
	}
}
