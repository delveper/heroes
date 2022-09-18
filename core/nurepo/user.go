package nurepo

import (
	"context"

	"github.com/delveper/heroes/core/ent"
)

type User = ent.User

func (kpr *Keeper) Add(usr User) (User, error) {
	SQL, err := genInsertQuery(usr)
	if err != nil {
		return User{}, err
	}
	kpr.conn.QueryRow(context.Background(), string(SQL))
	return User{}, nil
}
