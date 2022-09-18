package repo

import (
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"log"

	"github.com/delveper/heroes/cfg"

	_ "github.com/jackc/pgx/v4/stdlib" // we can switch drivers in config
	_ "github.com/lib/pq"
)

var (
	ErrInsertingValue      = errors.New("could not insert values into table")
	ErrDuplicateConstraint = errors.New("duplicate key value violates unique constraint")
	ErrEmailExists         = errors.New("email already exists") // smells bad
	errUnique              = errors.New("unique_violation")
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
	res, err := kpr.DB.Exec(userSQL)
	if err != nil {
		return fmt.Errorf("error has occured making migrations: %w", err)
	}
	log.Printf("migrations were made: %+v", res)
	return nil
}
