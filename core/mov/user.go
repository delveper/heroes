package mov

import (
	"errors"
	"net/http"

	"github.com/delveper/heroes/core/ent"
	"github.com/delveper/heroes/pkg/black"
	"github.com/lib/pq"
)

type UserMover interface { // we can use here repo.UserKeeper instead
	Add(ent.User) (ent.User, error)
}

var ErrUniqueEmail = errors.New("user email has to be unique")

// TODO: handle errors gracefully

func (mvr *Mover) HandleUser() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			mvr.Add(rw, req)
		case http.MethodGet:
		case http.MethodPut:
		case http.MethodDelete:
		default:
			respondHTTPErr(rw, req, http.StatusNotFound)
		}
	})
}

// Add is a method creating user
func (mvr *Mover) Add(rw http.ResponseWriter, req *http.Request) {
	// unmarshal user
	usr := ent.User{}
	if err := decodeBody(req, &usr); err != nil {
		respondErr(rw, req, http.StatusInternalServerError, err)
	}
	// validation
	switch err := black.ValidateStruct(usr); {
	case errors.As(err, new(*black.ValidationError)):
		respondErr(rw, req, http.StatusUnprocessableEntity, errors.Unwrap(err))
		return
	case err != nil:
		respondErr(rw, req, http.StatusInternalServerError, err)
		return
	default: // everything looks ok so far
	}
	// add usr TODO: Handle errors gracefully
	// switch looks more flat
	switch usr, err := mvr.Agent.Add(usr); {
	case err != nil:
		switch err, ok := err.(*pq.Error); ok {
		case err.Code.Name() == "unique_violation":
			respondErr(rw, req, http.StatusConflict, ErrUniqueEmail)
			return
		default:
			// predict other cases
		}
		respondErr(rw, req, http.StatusInternalServerError, err)
		return
	default:
		respond(rw, req, http.StatusCreated, usr)
	}
}
