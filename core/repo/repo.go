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

type Keeper struct {
	repo *sql.DB
}

// NewKeeper returns pointer receiver
// to new Peeper with database options
// from out the box
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

	return &Keeper{repo: db}, nil
}

// Add creates new record in database table user
func (k *Keeper) Add(usr ent.User) (ent.User, error) {
	SQL := `INSERT INTO "user" (full_name, email, password)
				VALUES($1, $2, crypt($3, gen_salt('md5')))
					RETURNING id,  created_at;`

	if err := k.repo.QueryRow(SQL, usr.FullName, usr.Email, usr.Password).
		Scan(&usr.ID, &usr.CreatedAt); err != nil {
		log.Printf("error occured inserting %+v: %v", usr, err)
		return ent.User{}, ErrInsertingValue
	}

	return usr, nil
}

// CreateTable will create table using given DDL query
func (k *Keeper) CreateTable(q string) error {
	res, err := k.repo.Exec(q)
	if err != nil {
		return err
	}
	log.Printf("table created successfully: %v", res)

	return nil
}
