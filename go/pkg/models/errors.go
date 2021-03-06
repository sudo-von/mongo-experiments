package models

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

// ErrResponse renderer type for handling all sorts of errors.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	ErrorCode int    `form:"error_code" json:"error_code" yaml:"error_code" xml:"error_code"`
	Message   string `form:"message" json:"message" yaml:"message" xml:"message"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		ErrorCode:      400,
		Message:        err.Error(),
	}
}

func ErrInternalServer(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		ErrorCode:      500,
		Message:        err.Error(),
	}
}

func ErrNotFound(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 404,
		ErrorCode:      404,
		Message:        err.Error(),
	}
}

func ErrUnauthorized(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 401,
		ErrorCode:      401,
		Message:        err.Error(),
	}
}

func ErrConflict(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 409,
		ErrorCode:      409,
		Message:        err.Error(),
	}
}

func ErrForbidden(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 403,
		ErrorCode:      403,
		Message:        err.Error(),
	}
}

func mergeErrors(e1, e2 error) error {
	if e1 != nil && e2 != nil {
		return errors.New(e1.Error() + "; " + e2.Error())
	}
	if e1 != nil {
		return e1
	}

	return e2
}
