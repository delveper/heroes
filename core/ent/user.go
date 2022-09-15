package ent

import (
	_ "embed"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

// User is a key entity
type User struct {
	ID        string    `json:"id"` //  uuid.UUID
	FullName  string    `json:"full_name" regex:"^[\p{L}a-zA-Z&\s-'’.]{2,255}$"`
	Email     string    `json:"email" regex:"^[_A-Za-z0-9-\+]+(\.[_A-Za-z0-9-]+)*@[A-Za-z0-9-]+(\.[A-Za-z0-9]+)*(\.[A-Za-z]{2,})$"`
	Password  string    `json:"password" regex:"^.{8,255}$" dfdfdf:"dfsfsaf"`
	CreatedAt time.Time `json:"created_at"`
}

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
func (serv *Service) CreateTable() error {
	return serv.Keeper.CreateTable()
}

// Add will add new user to database
// name of method might be changed
func (serv *Service) Add(usr User) (User, error) {
	// we will handle validation on transport layer

	// if err := usr.Validate(); err != nil {
	// 	log.Printf("error occurred validating %+v: %v", usr, err)
	// 	return User{}, err
	// }

	usr, err := serv.Keeper.Add(usr)
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
