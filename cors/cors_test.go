package cors

import (
	"testing"
)

func TestNewCORSOptions(t *testing.T) {

	cors := New(Options{
		MaxAge:           86400,
		AllowedHeaders:   []string{"X-Service-Role", "Authorization"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		AllowedOrigins:   []string{"*"},
	})

	cors.AddHeader("X-Test-Role")

	cors.AddHeader("X-Admin-Role")

	if cors.options.AllowedHeaders == nil {
		t.Error("Expected AllowedHeaders to be non-nil but got, nil instead")
	}

	var foundXAdminRole bool = false

	for _, v := range cors.options.AllowedHeaders {
		if v == "X-Admin-Role" {
			foundXAdminRole = true
		}
	}
	if !foundXAdminRole {
		t.Error("Expected AllowedHeaders to contain X-Admin-Role but could not find X-Admin-Role header")
	}
}

func TestDefaultCORSOptions(t *testing.T) {
	cors := Default()

	cors.AddHeader("X-Hello-Header")

	if	cors.options.AllowedHeaders == nil {
		t.Error("Expected AllowedHeaders to be non-nil but got, nil instead")
	}

	var foundHelloHeader bool = false

	for _, v := range cors.options.AllowedHeaders {
		if v == "X-Hello-Header" {
			foundHelloHeader = true
		}
	}
	if !foundHelloHeader {
		t.Error("Expected AllowedHeaders to contain X-Hello-Header but could not find X-Hello-Header header")
	}
}
