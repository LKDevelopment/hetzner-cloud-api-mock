package model

type APIRoute struct {
	Method       string
	Description  string
	Route        string
	Response     string
	ResponseCode int
	JsonSchema   string
}

type APIHandler struct {
	Name   string
	Routes []APIRoute
}
