package nurepo

import (
	"context"
	"fmt"

	"github.com/delveper/heroes/core/ent"
)

type User = ent.User

func (kpr *Keeper) Add(usr User) (User, error) {
	SQL, err := GenInsertQuery(usr)
	fmt.Println(string(SQL))
	if err != nil {
		return User{}, err
	}
	var id string
	if err := kpr.conn.QueryRow(context.Background(), string(SQL)).Scan(&id); err != nil {
		return User{}, err
	}

	return User{ID: id}, nil
}
