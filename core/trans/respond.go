package trans

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func decodeBody(req *http.Request, val any) error {
	defer req.Body.Close()
	return json.NewDecoder(req.Body).Decode(val)
}

func encodeBody(rw http.ResponseWriter, req *http.Request, val any) error {
	return json.NewEncoder(rw).Encode(val)
}

func respond(rw http.ResponseWriter, req *http.Request, status int, data any) {
	rw.WriteHeader(status)
	if data != nil {
		if err := encodeBody(rw, req, data); err != nil {
			respondErr(rw, req, http.StatusInternalServerError)
		}
	}
}

func respondErr(rw http.ResponseWriter, req *http.Request, status int, args ...interface{}) {
	respond(rw, req, status, map[string]interface{}{
		"error": map[string]interface{}{
			"message": fmt.Sprint(args...),
		},
	})
}

func respondHTTPErr(rw http.ResponseWriter, req *http.Request, status int) {
	respondErr(rw, req, status, http.StatusText(status))
}
