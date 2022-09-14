package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/delveper/heroes/cfg"
	"github.com/delveper/heroes/core/ent"
	_ "github.com/lib/pq"
)

var (
	ErrCreatingTable  = errors.New("could not create table")
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

// Add creates new record in database table user
func (kpr *Keeper) Add(usr ent.User) (ent.User, error) {
	SQL := `INSERT INTO "user" (full_name, email, password)
				VALUES($1, $2, crypt($3, gen_salt('md5')))
					RETURNING id,  created_at;`

	if err := kpr.QueryRow(SQL, usr.FullName, usr.Email, usr.Password).
		Scan(&usr.ID, &usr.CreatedAt); err != nil {
		log.Printf("error occured inserting %+v: %v", usr, err)
		return ent.User{}, fmt.Errorf("%v; %w", ErrInsertingValue, err)
	}

	return usr, nil
}

// CreateTable will create table using given DDL query
func (kpr *Keeper) CreateTable(q string) error {
	_, err := kpr.Exec(q)
	if err != nil {
		return err
	}
	return nil
}
