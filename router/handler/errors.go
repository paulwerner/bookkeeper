package handler

import (
	"net/http"

	d "github.com/paulwerner/bookkeeper/pkg/domain"
)

type errorResponse struct {
	Errors map[string]any `json:"errors"`
}

func newErrorResponse(err any) (errorResponse, int) {
	e := errorResponse{}
	e.Errors = make(map[string]any)

	switch serr := err.(type) {
	case d.NotFoundError:
		e.Errors["msg"] = serr.Error()
		return e, http.StatusNotFound

	case d.InvalidEntityError:
		e.Errors["msg"] = serr.Error()
		return e, http.StatusUnprocessableEntity

	case d.InvalidPasswordError:
		e.Errors["msg"] = serr.Error()
		return e, http.StatusForbidden

	case d.InvalidLengthError:
		e.Errors["msg"] = serr.Error()
		return e, http.StatusUnprocessableEntity

	case d.AlreadyInUseError:
		e.Errors["msg"] = serr.Error()
		return e, http.StatusConflict

	case d.InvalidAccessTokenError:
		e.Errors["msg"] = serr.Error()
		return e, http.StatusForbidden

	default:
		e.Errors["msg"] = "unexpected error occurred"
		return e, http.StatusInternalServerError
	}
}
