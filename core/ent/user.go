package ent

import (
	_ "embed"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

var (
	ErrCreatingUser = errors.New("could not create user")
	ErrInvalidUser  = errors.New("user has to be valid")
)

// DDL for user table
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
// about interfaces and structs
func NewService(uk UserKeeper) *Service {
	return &Service{Keeper: uk}
}

// SetID will set new UUID
// in case we would not provide one
// database will gen it automatically
func (usr *User) SetID() {
	usr.ID = uuid.New().String()
}

// Validate ensures us that User is OK
// otherwise returns error
// TODO: feels like something is missing
func (usr *User) Validate() error {
	if IsValidPassword(usr.Password) && IsValidEmail(usr.Email) && IsValidName(usr.FullName) {
		return nil // user is ok
	}
	return ErrInvalidUser
}

// CreateTable will create table using
// embedded SQL that stored in createTableSQL variable
// is a better choice to hardcode DDL?
func (srv *Service) CreateTable() error {
	return srv.Keeper.CreateTable(createTableSQL)
}

// Add will add new user to repo
// name of method might be changed
func (srv *Service) Add(usr User) (User, error) {
	if err := usr.Validate(); err != nil {
		log.Printf("error occured creating validating %+v: %v", usr, err)
		return User{}, ErrInvalidUser
	}

	usr, err := srv.Keeper.Add(usr)
	if err != nil {
		log.Printf("error occured creating user: %v", err)
		return User{}, ErrCreatingUser
	}

	return usr, nil
}
