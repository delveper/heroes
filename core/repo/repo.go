package repo

import (
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"log"

	"github.com/delveper/heroes/cfg"
	"github.com/delveper/heroes/core/ent"
	_ "github.com/lib/pq"
)

var (
	ErrInsertingValue = errors.New("could not insert values into table")
)

type Keeper struct{ *sql.DB }

// NewKeeper returns pointer receiver
// to new Keeper with database options out of the box
func NewKeeper(opt *cfg.Options) (*Keeper, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		opt.Repo.Host, opt.Repo.Port, opt.Repo.UserName, opt.Repo.Password, opt.Repo.DbName)

	db, err := sql.Open(opt.Repo.DriverName, dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Keeper{db}, nil
}

//go:embed sql/user.sql
var userSQL string

// MakeMigrations will create table using given DDL query
// un out particular case only user table
func (kpr *Keeper) MakeMigrations() error {
	res, err := kpr.Exec(userSQL)
	if err != nil {
		return fmt.Errorf("error has occured making migrations: %w", err)
	}
	log.Printf("migrations were made: %+v", res)
	return nil
}

// Add creates new record in database table user
func (kpr *Keeper) Add(usr ent.User) (ent.User, error) {
	SQL := `INSERT INTO "user" (full_name, email, password)
				VALUES($1, $2, crypt($3, gen_salt('md5')))
					RETURNING id,  created_at;`

	if err := kpr.QueryRow(SQL, usr.FullName, usr.Email, usr.Password).
		Scan(&usr.ID, &usr.CreatedAt); err != nil {
		// TODO: Handle errors gracefully
		log.Printf("error occured inserting %+v: %v", usr, err)
		return ent.User{}, fmt.Errorf("%v; %w", ErrInsertingValue, err)
	}

	return usr, nil
}
