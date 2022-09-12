package trans

import (
	"fmt"
	"net/http"

	"github.com/delveper/heroes/core/ent"
)

type UserService interface { // we can use here repo.UserKeeper instead
	Add(ent.User) (ent.User, error)
}

// TODO: handle errors gracefully

func UserHandler(hdl *Handler) http.Handler {
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
		respondHTTPErr(rw, req, http.StatusInternalServerError)
		return
	}
	fmt.Printf("%+v", usr)
	usr, err := hdl.Service.Add(usr)
	if err != nil {
		respondHTTPErr(rw, req, http.StatusBadRequest)
		return
	}

	respond(rw, req, http.StatusCreated, usr)
}
