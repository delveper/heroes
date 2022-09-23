package repo

import (
	"github.com/delveper/heroes/core/ent"
)

type User = ent.User // gives us ent.UserKeeper signature down to the last penny

// Add creates new record in database table user
func (kpr *Keeper) Add(usr User) (User, error) {
	SQL := `INSERT INTO "user" (first_name, last_name, email, password)
						 VALUES($1, $2, $3, crypt($4, gen_salt('md5')))
						 ON CONFLICT (id) DO UPDATE -- almost impossible  
						     SET id = gen_random_uuid() 
			RETURNING id,  created_at;`

	args := []interface{}{usr.FirstName, usr.LastName, usr.Email, usr.Password}
	dest := []interface{}{&usr.ID, &usr.CreatedAt}

	if err := kpr.QueryRow(SQL, args...).Scan(dest...); err != nil {
		return User{}, err
	}

	return usr, nil
}
