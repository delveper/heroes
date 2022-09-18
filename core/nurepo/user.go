package nurepo

import (
	"context"
	"fmt"

	"github.com/delveper/heroes/core/ent"
)

type User = ent.User

// Add implements UserKeeper TODO: Make methods type agnostic func(kpr *Keeper) Add(any) (any, error)
func (kpr *Keeper) Add(usr User) (User, error) {
	SQL, err := GenInsertQuery(usr)
	if err != nil {
		return User{}, err
	}

	var id interface{}
	if err := kpr.conn.QueryRow(context.Background(), string(SQL)).Scan(&id); err != nil {
		return User{}, err
	}

	fmt.Println(id)
	return User{}, nil
}

func (kpr *Keeper) Get(id string) (User, error) {
	SQL, err := GenInsertQuery(User{})
	if err != nil {
		return User{}, err
	}

	usr := User{}
	if err := kpr.conn.QueryRow(context.Background(), string(SQL)).Scan(&usr); err != nil {
		return User{}, err
	}
	return usr, nil
}
