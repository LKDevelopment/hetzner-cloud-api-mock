package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/lkdevelopment/hetzner-cloud-api-mock"
	"github.com/lkdevelopment/hetzner-cloud-api-mock/validator"

	"github.com/gorilla/mux"
	"github.com/lkdevelopment/hetzner-cloud-api-mock/model"
)

type Server struct {
	Handler     map[string]model.APIHandler
	Port        string
	APIFilePath string
	DebugMode   bool
}

type MockServerHealthResponse struct {
	MockServerVersion string `json:"mock_server_version"`
	APIVersion        string `json:"api_version"`
}

func (s *Server) Start() {
	r := mux.NewRouter()
	log.Printf("---------------Setup--------------\n")
	r.HandleFunc("/_info", s.infoHandler).Name("version").Methods(http.MethodGet)
	s.registerMockedRoutes(r)
	log.Printf("----------------End---------------\n")
	http.Handle("/", r)
	_ = http.ListenAndServe(":"+s.Port, nil)
}

func (s *Server) infoHandler(writer http.ResponseWriter, request *http.Request) {
	// get last modified time
	file, err := os.Stat(s.APIFilePath)
	if err != nil {
		log.Fatal(err)
	}
	apiFileModifiedTime := file.ModTime()
	writer.WriteHeader(http.StatusOK)
	mockServerHealth := MockServerHealthResponse{
		MockServerVersion: api_mock.Version,
		APIVersion:        apiFileModifiedTime.Format(time.RFC1123),
	}
	resp, err := json.Marshal(mockServerHealth)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = writer.Write(resp)
}

func (s *Server) registerMockedRoutes(r *mux.Router) {
	for _, handler := range s.Handler {
		log.Printf("RESOURCE %s\n", handler.Name)
		for _, apiRoutePointer := range handler.Routes {
			apiRoute := apiRoutePointer
			routeWithPrefix := strings.ReplaceAll(strings.ReplaceAll("/v1"+apiRoute.Route, "{id}", "{id:[0-9]+}"), "{action_id}", "{action_id:[0-9]+}")
			log.Printf("\t [%s] %s\n", apiRoute.Method, routeWithPrefix)
			r.HandleFunc(routeWithPrefix, func(writer http.ResponseWriter, request *http.Request) {
				ValidateRequestResponse, err := validator.ValidateRequest(request)
				if ValidateRequestResponse {
					body, _ := ioutil.ReadAll(request.Body)
					validateJsonSchemaResponse, err := validator.ValidateJsonSchema(string(body), apiRoute)
					if validateJsonSchemaResponse {
						vars := mux.Vars(request)
						if s.DebugMode {
							log.Printf("Got %s route, with parameters %v\n", apiRoute.Description, vars)
						}
						s.respondWith(writer, apiRoute.ResponseCode, []byte(apiRoute.Response))
					} else {
						resp, err := json.Marshal(err)
						if err != nil {
							log.Println(err)
						}
						s.respondWith(writer, http.StatusUnprocessableEntity, resp)
					}
				} else {
					statusCode := http.StatusUnauthorized
					if err.Code == "json_error" {
						statusCode = http.StatusBadRequest
					}
					resp, err := json.Marshal(err)
					if err != nil {
						log.Println(err)
					}
					s.respondWith(writer, statusCode, resp)
				}
			}).Name(apiRoute.Description).Methods(apiRoute.Method)
		}
	}
}

func (s *Server) respondWith(writer http.ResponseWriter, statusCode int, body []byte) {
	writer.WriteHeader(statusCode)
	_, err := writer.Write(body)
	if err != nil {
		log.Panic(err)
	}
}
