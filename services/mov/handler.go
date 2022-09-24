// Package mover is responsible for logistics

package mov

import (
	"github.com/delveper/heroes/core"
)

type (
	User      = core.User
	UserAgent = core.UserAgent
)

// Handler takes care of the http requests,
// responses and validations
type Handler struct {
	Agent UserAgent
}

// NewUserHandler  will accept interface UserAgent
// and return pointer receiver to Handler struct
func NewUserHandler(ua UserAgent) *Handler {
	return &Handler{ua}
}

func (mvr *Handler) Add(usr User) error {
	return nil
}
