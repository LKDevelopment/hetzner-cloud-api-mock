package model

import (
	"testing"

	"github.com/lkdevelopment/hetzner-cloud-api-mock/apiblueprint"
)

func TestTransformer_Transform(t *testing.T) {
	apibTransformer := apiblueprint.Transformer{
		Path:                 "../../assets/",
		ApiBlueprintFileName: "test.apib",
		JSONFileName:         "drafterOutput.json",
	}
	apibTransformer.TransformAPIBlueprintFile()
	transformer := Transformer{}
	transformed := transformer.Transform(apibTransformer.Resources)
	if len(transformed) != 2 {
		t.Errorf("Expected 2, got %v", len(transformed))
	}
	if _, ok := transformed["actions"]; !ok {
		t.Error("Expected true, got false")
	}
	if transformed["actions"].Name != "/actions" {
		t.Errorf("Expected /actions, got %v", transformed["actions"].Name)
	}
	if transformed["servers"].Name != "/servers" {
		t.Errorf("Expected /servers, got %v", transformed["servers"].Name)
	}
	if len(transformed["actions"].Routes) != 2 {
		t.Errorf("Expected 2, got %v", len(transformed["actions"].Routes))
	}
	route := transformed["actions"].Routes[0]
	if route.Method != "GET" {
		t.Errorf("Expected GET, got %v", route.Method)
	}
	if route.Description != "List all Actions" {
		t.Errorf("Expected \"List all Actions\", got %v", route.Description)
	}
	if route.Route != "/actions" {
		t.Errorf("Expected /actions, got %v", route.Route)
	}
	if route.Response == "" {
		t.Error("Expected non empty string, got empty string")
	}
	if route.JsonSchema != "" {
		t.Errorf("Expected empty string, got %v", route.JsonSchema)
	}

	routeServer := transformed["servers"].Routes[10]
	if routeServer.Method != "POST" {
		t.Errorf("Expected POST, got %v", routeServer.Method)
	}
	if routeServer.Description != "Enable Rescue Mode for a Server" {
		t.Errorf("Expected \"Enable Rescue Mode for a Server\", got %v", routeServer.Description)
	}
	if routeServer.Route != "/servers/{id}/actions/enable_rescue" {
		t.Errorf("Expected \"/servers/{id}/actions/enable_rescue\", got %v", routeServer.Route)
	}
	if routeServer.Response == "" {
		t.Error("Expected non empty string, got empty string")
	}
	if routeServer.JsonSchema == "" {
		t.Error("Expected non empty string, got empty string")
	}

}
