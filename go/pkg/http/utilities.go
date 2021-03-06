package http

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/mongo-experiments/go/pkg/models"
	"github.com/mongo-experiments/go/pkg/newerrors"

	"github.com/go-chi/render"
)

// Converts string to int.
func ParamToInt(param string, val url.Values) *int {
	if s := val.Get(param); s != "" {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil
		}
		return &i
	}
	return nil
}

// CheckError handles all the http error status response.
func CheckError(err error, w http.ResponseWriter, r *http.Request) {
	if _, ok := err.(*newerrors.ErrNotFound); ok {
		render.Render(w, r, models.ErrNotFound(err))

	} else if _, ok := err.(*newerrors.ErrBadRequest); ok {
		render.Render(w, r, models.ErrInvalidRequest(err))

	} else if _, ok := err.(*newerrors.ErrUnauthorized); ok {
		render.Render(w, r, models.ErrUnauthorized(err))

	} else if _, ok := err.(*newerrors.ErrForbidden); ok {
		render.Render(w, r, models.ErrForbidden(err))
	} else if _, ok := err.(*newerrors.ErrConflict); ok {
		render.Render(w, r, models.ErrConflict(err))

	} else {
		render.Render(w, r, models.ErrInternalServer(err))
	}
}
