package lib

import "net/http"

// ApplicationException - application error messages and codes
type ApplicationException struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HTTPStatus int    `json:"-"`
}

// JSONParseError - occurs when cannot marshal json object
var JSONParseError = ApplicationException{
	Code:       "JSONParseError",
	Message:    "Something went wrong",
	HTTPStatus: http.StatusInternalServerError,
}
