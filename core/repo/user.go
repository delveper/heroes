package repo

import (
	"fmt"
	"strings"

	"github.com/delveper/heroes/core/ent"
)

// Add creates new record in database table user
func (kpr *Keeper) Add(usr ent.User) (ent.User, error) {
	SQL := `INSERT INTO "user" (full_name, email, password)
				VALUES($1, $2, crypt($3, gen_salt('md5')))
					RETURNING id,  created_at;`

	//  TODO: Handle errors gracefully
	switch err := kpr.QueryRow(SQL, usr.FullName, usr.Email, usr.Password).
		Scan(&usr.ID, &usr.CreatedAt); {

	case err == nil: // return

	case strings.Contains(err.Error(), ErrDuplicateConstraint.Error()):
		return ent.User{}, fmt.Errorf("%v: %w", ErrInsertingValue, ErrEmailExists)

	default:
		return ent.User{}, fmt.Errorf("%v: %w", ErrInsertingValue, err)

	}
	return usr, nil
}
