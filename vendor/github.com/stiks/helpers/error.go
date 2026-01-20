package helpers

import (
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
)

type ValidationErr struct {
	Errors interface{} `json:"errors"`
}

// NewValidationError creates a new HTTPError instance.
func NewValidationError(errs error) *echo.HTTPError {
	if e, ok := errs.(validation.InternalError); ok {
		// an internal error happened
		fmt.Println(e.InternalError())
	}

	return &echo.HTTPError{Code: http.StatusBadRequest, Message: ValidationErr{Errors: errs}}
}
