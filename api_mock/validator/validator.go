package validator

import (
	"fmt"
	"io/ioutil"

	"net/http"
	"regexp"

	"github.com/lkdevelopment/hetzner-cloud-api-mock/model"
	"github.com/xeipuuv/gojsonschema"
)

func ValidateJsonSchema(body string, route model.APIRoute) (bool, *model.ErrorResponse) {
	if len(body) > 0 {
		schemaLoader := gojsonschema.NewStringLoader(route.JsonSchema)
		documentLoader := gojsonschema.NewStringLoader(body)
		result, err := gojsonschema.Validate(schemaLoader, documentLoader)
		if err != nil {
			panic(err.Error())
		}
		if result.Valid() {
			return true, nil
		} else {
			req, _ := regexp.Compile(`(?m)(\b[^\s]+\b)`) // Split every word from the error message
			schemaErr := model.ErrorResponse{
				Code: "invalid_input",
			}
			for _, desc := range result.Errors() {
				errorMessage := req.FindAllString(desc.String(), -1)
				field := errorMessage[0]
				if field == "root" { // when the field is root, te second parameter is the field that will be required
					// errorMessage := [root, image, is, required]
					field = errorMessage[1]
					schemaErr.Message = fmt.Sprintf("invalid input in field '%s'", field)
					schemaErr.Details = model.ErrorDetailsInvalidInput{
						Fields: []model.ErrorDetailsInvalidInputField{
							{
								Name: field,
								Messages: []string{
									"Missing data for required field.",
								},
							},
						},
					}
				} else {
					// errorMessage := [upgrade_disk, Invalid, type, Expected, boolean, given, string]
					// or
					// errorMessage := [type, must, be, one, of, the, following, linux64, linux32, freebsd64]
					if errorMessage[1] == "Invalid" {
						schemaErr.Message = fmt.Sprintf("invalid input in field '%s'", field)
						schemaErr.Details = model.ErrorDetailsInvalidInput{
							Fields: []model.ErrorDetailsInvalidInputField{
								{
									Name: field,
									Messages: []string{
										model.GetInValidTypeErrorMessage(errorMessage[len(errorMessage)-3]),
									},
								},
							},
						}
					} else if errorMessage[2] == "must" {
						schemaErr.Message = fmt.Sprintf("invalid input in field '%s'", field)
						schemaErr.Details = model.ErrorDetailsInvalidInput{
							Fields: []model.ErrorDetailsInvalidInputField{
								{
									Name: field,
									Messages: []string{
										"Not a valid choice.",
									},
								},
							},
						}
					}
				}
			}
			return false, &schemaErr
		}
	}
	return true, nil
}

func ValidateRequest(r *http.Request) (bool, *model.ErrorResponse) {
	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		return false, &model.ErrorResponse{
			Code:    "unauthorized",
			Message: "unable to authenticate",
		}
	}
	if r.Body != nil {
		body, _ := ioutil.ReadAll(r.Body)
		if len(body) > 0 {
			contentType := r.Header.Get("Content-Type")
			if contentType != "application/json" {
				return false, &model.ErrorResponse{
					Code:    "json_error",
					Message: "Content-Type must be set to application/json",
				}

			}
		}
	}
	return true, nil

}
