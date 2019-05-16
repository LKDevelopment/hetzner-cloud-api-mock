package model

import (
	"regexp"
	"strings"

	"github.com/bukalapak/snowboard/api"
)

type Transformer struct{}

func (p *Transformer) Transform(resources []*api.Resource) map[string]APIHandler {
	return p.groupIntoAPIHandler(p.mapIntoApiRoutes(resources))
}

func (p *Transformer) mapIntoApiRoutes(resources []*api.Resource) (resp []APIRoute) {
	re := regexp.MustCompile(`\{\?[^{\?}]*\}`) // Remove all GET Parameters TODO: We should support the GET Parameters too
	for _, resource := range resources {
		for _, method := range resource.Transitions {
			resp = append(resp, APIRoute{
				Method:       method.Method,
				Description:  method.Title,
				Route:        re.ReplaceAllString(method.Href.Path, ""),
				Response:     method.Transactions[0].Response.Body.Body,
				ResponseCode: method.Transactions[0].Response.StatusCode,
				JsonSchema:   method.Transactions[0].Request.Schema.Body,
			})
		}
	}
	return
}

func (p *Transformer) groupIntoAPIHandler(routes []APIRoute) map[string]APIHandler {
	handler := make(map[string]APIHandler)
	for _, route := range routes {
		resourceParts := strings.Split(route.Route, "/")
		index := resourceParts[1]
		if _, ok := handler[index]; !ok {
			handler[index] = APIHandler{
				Name:   route.Route,
				Routes: []APIRoute{},
			}
		}
		_handler := handler[index]
		_handler.Routes = append(_handler.Routes, route)
		handler[index] = _handler
	}
	return handler
}
