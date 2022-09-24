// Package keeper will act for CRUD operations with database.
//

package repo

import (
	"database/sql"

	"github.com/delveper/heroes/core"
)

type (
	User = core.User
	// UserKeeper = core.UserKeeper
)

type Keeper struct {
	db *sql.DB
}

// NewKeeper will initialise database
func NewKeeper(db *sql.DB) *Keeper {
	return &Keeper{db}
}

func (kpr *Keeper) Add(User) error {
	return nil
}
