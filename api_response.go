package goutyl

import (
	"net/http"
)

type ApiResponse struct {
	Error interface{} `json:"error,omitempty"`
	statusCode int
	Message string `json:"message,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}


func (r ApiResponse) JsonString() string {
	jsonString, err := JsonString(r)

	if err != nil {
		panic(err)
	}
	return jsonString
}

func (r ApiResponse) Json(w http.ResponseWriter) {
	if r.statusCode == 0 {
		r.statusCode = 200
	}

	jsonByteArr, err := Json(r)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(r.statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonByteArr)
}

func (r *ApiResponse) Status(code int) *ApiResponse {
	r.statusCode = code
	return r
}