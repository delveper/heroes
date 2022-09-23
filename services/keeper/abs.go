// Package keeper will act for CRUD to database.
// Using abstract interfaces without specific
// knowledge of the implementation details
// will make our software flexible and maintainable
package keeper

import (
	"github.com/delveper/heroes/core/ent"
)

type User = ent.User

type UserKeeper interface {
	Add(User) error
}
