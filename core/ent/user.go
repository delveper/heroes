package ent

import (
	_ "embed"
	"strings"
	"time"
)

// User is a key entity
// `regex` was designed for field validation ./pkg/black
type User struct {
	ID        string    `json:"id"` // may be uuid.UUID
	FirstName string    `json:"first_name" regex:"(?i)^[\p{L}A-Z&\s-'’.]{2,255}$"`
	LastName  string    `json:"last_name" regex:"(?i)^[\p{L}A-Z&\s-'’.]{2,255}$"`
	Email     string    `json:"email" regex:"(?i)^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$"`
	Password  string    `json:"password" regex:"^.{8,255}$"`
	CreatedAt time.Time `json:"created_at"`
}

// UserKeeper defines an interface
// we want our logic to implement
type UserKeeper interface {
	Add(User) (User, error)
}

// Agent 006 will be user servant :)
type Agent struct {
	Keeper UserKeeper
}

// NewAgent is proverbial case
// about passing interfaces and returning structs
func NewAgent(uk UserKeeper) *Agent {
	return &Agent{Keeper: uk}
}

// Clean will make our awesome user like an angel
func (usr *User) Clean() {
	usr.FirstName = strings.TrimSpace(usr.FirstName)
	usr.LastName = strings.TrimSpace(usr.LastName)
	usr.Email = strings.TrimSpace(usr.Email)
}

// Add will create new user and add it to database
func (agt *Agent) Add(usr User) (User, error) {
	return agt.Keeper.Add(usr)
}
