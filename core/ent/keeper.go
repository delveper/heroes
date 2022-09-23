package ent

type Keeper interface {
	Add(User) error
	// ...
}

// Validator is self-explanatory
type Validator interface {
	Check(User) error
}
