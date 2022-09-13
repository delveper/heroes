package trans

import (
	"net/http"

	"github.com/delveper/heroes/core/ent"
)

type UserService interface { // we can use here repo.UserKeeper instead
	Add(ent.User) (ent.User, error)
}

// TODO: handle errors gracefully

func (hdl *Handler) UserHandle() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			hdl.Add(rw, req)
		case http.MethodGet:
		case http.MethodPut:
		case http.MethodDelete:
		default:
			respondHTTPErr(rw, req, http.StatusNotFound)
		}
	})
}

// Add is a method creating user
func (hdl *Handler) Add(rw http.ResponseWriter, req *http.Request) {
	usr := ent.User{}
	if err := decodeBody(req, &usr); err != nil {
		respondErr(rw, req, http.StatusInternalServerError, err)
		return
	}

	usr, err := hdl.Service.Add(usr)
	switch {
	case err == ent.ErrEmailExists:
		respondErr(rw, req, http.StatusConflict, err)
	case err != nil:
		respondErr(rw, req, http.StatusUnprocessableEntity, err)
	default:
		respond(rw, req, http.StatusCreated, usr)
	}
}
