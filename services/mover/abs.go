package mover

import (
	"github.com/delveper/heroes/core/ent"
)

type User = ent.User

type UserMover interface {
	Add(User) error
}
