package sply

import (
	"errors"
	"log"
	"net/http"

	"github.com/delveper/heroes/core/ent"
	"github.com/delveper/heroes/core/repo"
	"github.com/delveper/heroes/pkg/black"
)

type UserMover interface { // we can use here repo.UserKeeper instead
	Add(ent.User) (ent.User, error)
}

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
	usr := ent.User{}
	if err := decodeBody(req, &usr); err != nil {
		respondErr(rw, req, http.StatusInternalServerError, err)
		return
	}

	// validate usr
	switch err := black.ValidateStruct(usr); {

	case err == nil: // move forward

	case errors.As(err, new(*black.ValidationError)): // we do need pointer squared here
		log.Println(err)
		respondErr(rw, req, http.StatusUnprocessableEntity, errors.Unwrap(err))
		return

	default:
		respondErr(rw, req, http.StatusInternalServerError, err)
		return
	}

	// add usr
	// TODO: Handle errors gracefully
	switch usr, err := mvr.Agent.Add(usr); {

	case err == nil:
		respond(rw, req, http.StatusCreated, usr)

	case errors.Is(err, repo.ErrEmailExists): // how to decouple from repo and user?
		respondErr(rw, req, http.StatusConflict, errors.Unwrap(err))

	default:
		respondErr(rw, req, http.StatusInternalServerError, err)
	}
}
