package goutyl


import (
	"encoding/json"
)

func JsonString(obj interface{}) (string, error) {

	jsonString, err := Json(obj)

	if err != nil {
		return "", err
	}

	return string(jsonString[:]), nil
}

func Json(obj interface{}) ([]byte, error) {
	jsonByteArr, err := json.MarshalIndent(obj, "", "  ")

	if err != nil {
		return []byte(""), nil
	}

	return jsonByteArr, nil
}