package api

import (
	utyl "github.com/peek4y/goutyl"
	"net/http"
)

type Response struct {
	Error      interface{} `json:"error,omitempty"`
	statusCode int
	Message    string      `json:"message,omitempty"`
	Payload    interface{} `json:"payload,omitempty"`
}

func (r Response) JsonString() string {
	jsonString, err := utyl.JsonString(r)

	if err != nil {
		panic(err)
	}
	return jsonString
}

func (r Response) Json(w http.ResponseWriter) {
	if r.statusCode == 0 {
		r.statusCode = 200
	}

	jsonByteArr, err := utyl.Json(r)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(r.statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonByteArr)
}

func (r *Response) Status(code int) *Response {
	r.statusCode = code
	return r
}
