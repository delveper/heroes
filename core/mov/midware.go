package mov

import (
	"context"
	"errors"
	"net/http"

	"github.com/delveper/heroes/pkg/black"
)

var (
	ErrMissingParameter = errors.New("parameter is missing: ")
)

// Middleware is simple white magic
// for wrapping functional options
type Middleware func(http.Handler) http.Handler

type contextKey struct{ string }

func Wrap(hdl http.Handler, middleware ...Middleware) http.Handler {
	for _, mid := range middleware {
		hdl = mid(hdl)
	}
	return hdl
}

func WithJSON(hdl http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
		hdl.ServeHTTP(rw, req)
	})
}

// ValidateEntity will check if ent is OK
func ValidateEntity(src any) Middleware {
	return func(hdl http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			// unmarshal ent
			if err := decodeBody(req, src); err != nil {
				respondErr(rw, req, http.StatusInternalServerError, err)
			}
			// validate ent
			switch err := black.ValidateStruct(src); {
			case errors.As(err, new(*black.ValidationError)):
				respondErr(rw, req, http.StatusUnprocessableEntity, errors.Unwrap(err))
				return
			case err != nil:
				respondErr(rw, req, http.StatusInternalServerError, err)
				return
			default: // everything looks ok so far
				// struct underlying name
				name, err := black.GetStructName(src)
				if err != nil {
					respondErr(rw, req, http.StatusInternalServerError, err)
				}
				// just drilling techniques, maybe it is context violation
				ctx := context.WithValue(context.Background(), &contextKey{name}, src)
				hdl.ServeHTTP(rw, req.WithContext(ctx))
			}
		})
	}
}

// ValidateQueryParams check if all necessary parameters for
// main random logic are present in query
func ValidateQueryParams(params ...string) Middleware {
	return func(hdl http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			q := req.URL.Query()
			for _, param := range params {
				val := q.Get(param)
				if val == "" {
					respondErr(rw, req, http.StatusBadRequest, ErrMissingParameter, param)
					return
				}
			}
			hdl.ServeHTTP(rw, req) // all params are OK
		})
	}
}
