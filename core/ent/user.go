package ent

import (
	_ "embed"
	"errors"
	"strings"
	"time"
)

// User is a key entity
type User struct {
	ID        string    `json:"id"` // may be uuid.UUID
	FullName  string    `json:"full_name" regex:"^[\p{L}a-zA-Z&\s-'’.]{2,255}$"`
	Email     string    `json:"email" regex:"^[_A-Za-z0-9-\+]+(\.[_A-Za-z0-9-]+)*@[A-Za-z0-9-]+(\.[A-Za-z0-9]+)*(\.[A-Za-z]{2,})$"`
	Password  string    `json:"password" regex:"^.{8,255}$"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	ErrDuplicateConstraint = errors.New("duplicate key value violates unique constraint")
	ErrEmailExists         = errors.New("email already exists")
	ErrCreatingUser        = errors.New("could not create user")
)

// UserKeeper defines an interface
// we want our logic to implement
type UserKeeper interface {
	CreateTable() error
	Add(User) (User, error)
}

type Service struct {
	Keeper UserKeeper
}

// NewService is proverbial case
// about passing interfaces and returning structs
func NewService(uk UserKeeper) *Service {
	return &Service{Keeper: uk}
}

// TODO: feels like something is missing

// Clean will make our awesome user clean
func (usr *User) Clean() {
	usr.FullName = strings.TrimSpace(usr.FullName)
	usr.Email = strings.TrimSpace(usr.Email)
}

// CreateTable is what it is
func (serv *Service) CreateTable() error {
	return serv.Keeper.CreateTable()
}

// Add will create new user and add it to database
func (serv *Service) Add(usr User) (User, error) {
	usr, err := serv.Keeper.Add(usr)
	if err != nil {
		// feels stupid
		if strings.Contains(err.Error(), ErrDuplicateConstraint.Error()) {
			return User{}, ErrEmailExists
		}
		return User{}, ErrCreatingUser
	}

	return usr, nil
}
