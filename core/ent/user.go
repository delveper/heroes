package ent

import (
	_ "embed"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	ErrCreatingUser        = errors.New("could not create user")
	ErrInvalidEmail        = errors.New("user has to have valid email address")
	ErrEmailExists         = errors.New("user has to have unique email")
	ErrInvalidPassword     = errors.New("user has to have valid password")
	ErrInvalidName         = errors.New("user has to have valid name")
	ErrDuplicateConstraint = errors.New("duplicate key value violates unique constraint") // feels not OK
)

// not sure about that
// hard coding feels less appropriate
//
//go:embed user.sql
var createTableSQL string

// User is a key entity
type User struct {
	ID        string    `json:"id" sql:"id"` //  uuid.UUID
	FullName  string    `json:"full_name" sql:"full_name"`
	Email     string    `json:"email" sql:"email"`
	Password  string    `json:"password" sql:"password"`
	CreatedAt time.Time `json:"created_at" sql:"created_at"`
}

// UserKeeper defines an interface
// we want our logic to implement
type UserKeeper interface {
	CreateTable(string) error
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

// SetID will set new UUID
// in case we would not provide one
// database will gen it
func (usr *User) SetID() {
	usr.ID = uuid.New().String()
}

// Validate ensures us that User is OK otherwise returns error
// TODO: feels like something is missing
//
//	and I do not get how to deal with cascade of errors
func (usr *User) Validate() (err error) {
	// we want our awesome users name and email be clean
	// is it a right place to do that?
	usr.FullName = strings.TrimSpace(usr.FullName)
	usr.Email = strings.TrimSpace(usr.Email)

	switch {
	case !IsValidName(usr.FullName):
		return ErrInvalidName

	case !IsValidEmail(usr.Email):
		return ErrInvalidEmail

	case !IsValidPassword(usr.Password):
		return ErrInvalidPassword

	default:
		return nil
	}
}

// CreateTable will create table using
// embedded SQL that stored in createTableSQL variable
func (srv *Service) CreateTable() error {
	return srv.Keeper.CreateTable(createTableSQL)
}

// Add will add new user to database
// name of method might be changed
func (srv *Service) Add(usr User) (User, error) {
	if err := usr.Validate(); err != nil {
		log.Printf("error occured validating %+v: %v", usr, err)
		return User{}, err
	}

	usr, err := srv.Keeper.Add(usr)
	if err != nil {
		log.Printf("error occured creating user: %v", err)

		// feels stupid
		if strings.Contains(err.Error(), ErrDuplicateConstraint.Error()) {
			return User{}, ErrEmailExists
		}

		return User{}, ErrCreatingUser
	}

	return usr, nil
}
