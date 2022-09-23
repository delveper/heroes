// Package mover is responsible for logistics

package mover

import (
	"github.com/delveper/heroes/core"
)

type User = core.User

type UserMover interface {
	Add(User) error
}

// Mover  will accept input from client,
// preprocess it and send to promoter.Promoter
type Mover struct {
	UserMover
}

// NewUserMover  will accept interface UserMover
// and return pointer receiver to Mover struct
func NewUserMover(um UserMover) *Mover {
	return &Mover{um}
}
