package validator

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestValidateRequest(t *testing.T) {
	t.Run("Test missing Authorization", func(t *testing.T) {
		resp, err := ValidateRequest(&http.Request{})
		if resp {
			t.Error("Expected false, got true")
		}
		if err.Code != "unauthorized" {
			t.Errorf("Expected \"unauthorized\", got %v", err.Code)
		}
		if err.Message != "unable to authenticate" {
			t.Errorf("Expected \"unable to authenticate\", got %v", err.Message)
		}
		if err.Details != nil {
			t.Errorf("Expected nil, got %v", err.Details)
		}
	})
	t.Run("Test missing Content Type", func(t *testing.T) {
		resp, err := ValidateRequest(&http.Request{
			Header: map[string][]string{
				"Authorization": {"ABCD"},
			},
			Body: ioutil.NopCloser(strings.NewReader("{\"key\":\"value\"}")),
		})
		if resp {
			t.Error("Expected false, got true")
		}
		if err.Code != "json_error" {
			t.Errorf("Expected \"json_error\", got %v", err.Code)
		}
		if err.Message != "Content-Type must be set to application/json" {
			t.Errorf("Expected \"Content-Type must be set to application/json\", got %v", err.Message)
		}
		if err.Details != nil {
			t.Errorf("Expected nil, got %v", err.Details)
		}
	})
	t.Run("Test wrong Content Type", func(t *testing.T) {
		resp, err := ValidateRequest(&http.Request{
			Header: map[string][]string{
				"Authorization": {"ABCD"},
				"Content-Type":  {"application/not-json"},
			},
			Body: ioutil.NopCloser(strings.NewReader("{\"key\":\"value\"}")),
		})
		if resp {
			t.Error("Expected false, got true")
		}
		if err.Code != "json_error" {
			t.Errorf("Expected \"json_error\", got %v", err.Code)
		}
		if err.Message != "Content-Type must be set to application/json" {
			t.Errorf("Expected \"Content-Type must be set to application/json\", got %v", err.Message)
		}
		if err.Details != nil {
			t.Errorf("Expected nil, got %v", err.Details)
		}
	})
	t.Run("Test all OK", func(t *testing.T) {
		resp, err := ValidateRequest(&http.Request{
			Header: map[string][]string{
				"Authorization": {"ABCD"},
				"Content-Type":  {"application/json"},
			},
		})
		if !resp {
			t.Error("Expected true, got false")
		}
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	})
}
