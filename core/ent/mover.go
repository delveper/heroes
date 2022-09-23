package ent

type Mover interface {
	Keeper
	Validate(User) error
}
