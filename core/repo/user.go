package repo

import (
	"fmt"
	"strings"

	"github.com/delveper/heroes/core/ent"
)

type User = ent.User // gives us ent.UserKeeper signature down to the last penny

// Add creates new record in database table user
func (kpr *Keeper) Add(usr User) (User, error) {
	SQL := `INSERT INTO "user" (full_name, email, password)
				VALUES($1, $2, crypt($3, gen_salt('md5')))
					RETURNING id,  created_at;`

	args := []interface{}{usr.FullName, usr.Email, usr.Password}
	dest := []interface{}{&usr.ID, &usr.CreatedAt}

	switch err := kpr.QueryRow(SQL, args...).Scan(dest...); {
	//  TODO: Handle errors gracefully
	case err == nil: // return

	case strings.Contains(err.Error(), ErrDuplicateConstraint.Error()): // smells badly
		return ent.User{}, fmt.Errorf("%v: %w", ErrInsertingValue, ErrEmailExists)

	default:
		return ent.User{}, fmt.Errorf("%v: %w", ErrInsertingValue, err)
	}

	return usr, nil
}
