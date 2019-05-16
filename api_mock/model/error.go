package model

// ErrorResponse represents the schema of an error response.
type ErrorResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

// ErrorDetailsInvalidInput defines the schema of the Details field
// of an error with code 'invalid_input'.
type ErrorDetailsInvalidInput struct {
	Fields []ErrorDetailsInvalidInputField `json:"fields"`
}
type ErrorDetailsInvalidInputField struct {
	Name     string   `json:"name"`
	Messages []string `json:"messages"`
}

func GetInValidTypeErrorMessage(neededType string) string {
	respond := ""
	switch neededType {
	case "number":
		respond = "Not a valid integer."
		break
	case "string":
		respond = "Not a valid string."
		break
	case "boolean":
		respond = "Not a valid boolean."
		break
	}
	return respond
}
