// Package promoter holds on services
// that build all together the business flow
package promoter

import (
	"github.com/delveper/heroes/core"
)

type User = core.User

type UserPromoter interface {
	Add(User) error
}

type Promoter struct {
	UserPromoter
}

func NewUserPromoter(up UserPromoter) *Promoter {
	return &Promoter{up}
}
