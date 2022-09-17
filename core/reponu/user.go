package reponu

import (
	"context"
	"fmt"

	"github.com/delveper/heroes/core/ent"
)

type User = ent.User

func (kpr *Keeper) Add(usr User) (User, error) {
	var SQL string
	ctx := context.Background()

	SQL = `INSERT INTO "user" (full_name, email, password)
				VALUES($1, $2, crypt($3, gen_salt('md5')))
					RETURNING id,  created_at;`
	//  TODO: Handle errors gracefully
	switch err := kpr.conn.QueryRow(ctx, SQL, usr.FullName, usr.Email, usr.Password).
		Scan(&usr.ID, &usr.CreatedAt); {

	case err == nil: // return

	// case strings.Contains(err.Error(), ErrDuplicateConstraint.Error()):
	// 	return ent.User{}, fmt.Errorf("%v: %w", ErrInsertingValue, ErrEmailExists)

	default:
		return ent.User{}, fmt.Errorf("%v: %w", "", err)

	}

	return User{}, nil
}
