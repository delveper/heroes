package keeper

import (
	"database/sql"
)

type Keeper struct{ *sql.DB }

func NewKeeper(driver *sql.DB) *Keeper {
	return &Keeper{driver}
}

func (kpr *Keeper) Add(User) error {
	return nil
}
