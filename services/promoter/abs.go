package promoter

import (
	"github.com/delveper/heroes/core/ent"
)

type User = ent.User

type UserPromoter interface {
	Add(User) error
}
