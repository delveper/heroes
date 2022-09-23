package mov

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// we can cut it down to few lines
func decodeBody(req *http.Request, val any) error {
	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(val); err != nil {
		return fmt.Errorf("error decoding request body: %w", err)
	}

	return nil
}

func encodeBody(rw http.ResponseWriter, req *http.Request, val any) error {
	if err := json.NewEncoder(rw).Encode(val); err != nil {
		return fmt.Errorf("error encoding body: %w", err)
	}

	return nil
}

func respond(rw http.ResponseWriter, req *http.Request, sts int, data any) {
	rw.WriteHeader(sts)
	if data != nil {
		if err := encodeBody(rw, req, data); err != nil {
			respondErr(rw, req, http.StatusInternalServerError)
		}
	}
}

func respondErr(rw http.ResponseWriter, req *http.Request, sts int, args ...interface{}) {
	respond(rw, req, sts, map[string]interface{}{
		"error": map[string]interface{}{
			"message": fmt.Sprint(args...),
		},
	})
}

func respondHTTPErr(rw http.ResponseWriter, req *http.Request, sts int) {
	respondErr(rw, req, sts, http.StatusText(sts))
}
