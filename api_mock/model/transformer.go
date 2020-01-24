package model

import (
	"encoding/json"
	"log"
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
				Response:     p.addSampleMetaDataToResponse(method.Transactions[0].Response.Body.Body, method.Method, re.ReplaceAllString(method.Href.Path, "")),
				ResponseCode: method.Transactions[0].Response.StatusCode,
				JsonSchema:   method.Transactions[0].Request.Schema.Body,
			})
		}
	}
	return
}
func (p *Transformer) addSampleMetaDataToResponse(body string, method string, path string) string {
	if method == "GET" && strings.Count(path, "/") == 1 {
		var result map[string]interface{}
		json.Unmarshal([]byte(body), &result)
		for key, value := range result {
			// Each value is an interface{} type, that is type asserted as a string
			log.Println(key, value)
		}
		samplePagination := make(map[string]string)
		samplePagination["page"] = "1"
		samplePagination["per_page"] = "25"
		samplePagination["previous_page"] = "1"
		samplePagination["next_page"] = ""
		samplePagination["last_page"] = "1"
		samplePagination["total_entries"] = "1"
		meta := make(map[string]map[string]string)
		meta["pagination"] = samplePagination
		result["meta"] = meta
		data, _ := json.Marshal(result)
		return string(data)
	}
	return body
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
