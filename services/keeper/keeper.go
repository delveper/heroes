// Package keeper will act for CRUD to database.
// Using abstract interfaces without specific
// knowledge of the implementation details
// will make our software flexible and maintainable
// this pattern will follow on
package keeper

import (
	"database/sql"

	"github.com/delveper/heroes/core"
)

type User = core.User

type UserKeeper interface {
	Add(User) error
}

type Keeper struct{ *sql.DB }

func NewKeeper(db *sql.DB) *Keeper {
	return &Keeper{db}
}
