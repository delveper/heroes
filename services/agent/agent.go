// Package agent holds on services
// that build all together the business flow,
// agent only depends on package core.
package agent

import (
	"github.com/delveper/heroes/core"
)

type (
	User       = core.User
	UserKeeper = core.UserKeeper
	UserMover  = core.UserAgent
)

type Agent struct {
	Repository UserKeeper
}

func NewUserAgent(uk UserKeeper) *Agent {
	return &Agent{Repository: uk}
}

func (agt *Agent) Add(User) error {
	return nil
}

// type UserAgent struct {
// 	keeper UserKeeper
// }
//
// func NewUserAgent(uk UserKeeper) *UserAgent {
// 	return &UserAgent{keeper: uk}
// }
