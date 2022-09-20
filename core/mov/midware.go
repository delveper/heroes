package mov

import (
	"errors"
	"net/http"
)

var (
	ErrMissingParameter = errors.New("parameter is missing: ")
)

// Middleware is simple white magic
// for wrapping functional options
type Middleware func(http.Handler) http.Handler

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
