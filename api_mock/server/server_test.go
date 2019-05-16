package server

import (
	"net/http"
	"strings"
	"testing"

	"github.com/lkdevelopment/hetzner-cloud-api-mock/validator"
)

func TestValidateRequest(t *testing.T) {
	t.Run("everything ok", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/test", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header = http.Header{
			"Authorization": []string{"Bearer this-is-my-token"},
			"Content-Type":  []string{"application/json"},
		}
		resp, respErr := validator.ValidateRequest(req)
		if !resp {
			t.Error("Expected true, got false")
		}
		if respErr != nil {
			t.Error("Expected not error, got an error")
		}

	})
	t.Run("missing authorization header", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/test", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header = http.Header{
			"Content-Type": []string{"application/json"},
		}
		resp, respErr := validator.ValidateRequest(req)
		if resp {
			t.Error("Expected false, got true")
		}
		if respErr.Code != "unauthorized" {
			t.Errorf("Expected unauthorized, got %v", respErr.Code)
		}
		if !strings.Contains(respErr.Message, "unable to authenticate") {
			t.Error("Expected response to contain unable to authenticate")
		}
	})
	t.Run("missing content type", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/test", strings.NewReader("{\"key\":\"value\"}"))
		if err != nil {
			t.Fatal(err)
		}
		req.Header = http.Header{
			"Authorization": []string{"Bearer this-is-my-token"},
		}
		resp, respErr := validator.ValidateRequest(req)
		if resp {
			t.Error("Expected false, got true")
		}
		if respErr.Code != "json_error" {
			t.Errorf("Expected json_error, got %v", respErr.Code)
		}
		if !strings.Contains(respErr.Message, "Content-Type must be set to application/json") {
			t.Error("Expected response to contain Content-Type must be set to application/json")
		}
	})
}
