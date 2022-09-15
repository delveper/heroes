package ent

import (
	"errors"
	"regexp"
)

const (
	fullNamePattern = "^[\\p{L}a-zA-Z&\\s-'’.]{2,255}$"
	emailPattern    = "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	passwordPattern = "^.{8,255}$" // do we need to put any constraint on passwords except length?
	maxLength       = 255
)

var (
	ErrCreatingUser        = errors.New("could not create user")
	ErrInvalidEmail        = errors.New("user has to have valid email address")
	ErrEmailExists         = errors.New("user has to have unique email")
	ErrInvalidPassword     = errors.New("user has to have valid password")
	ErrInvalidName         = errors.New("user has to have valid name")
	ErrDuplicateConstraint = errors.New("duplicate key value violates unique constraint") // feels not OK
)

func IsValidEmail(str string) bool {
	return regexp.MustCompile(emailPattern).MatchString(str) && len(str) <= maxLength
}

func IsValidName(str string) bool {
	return regexp.MustCompile(fullNamePattern).MatchString(str)
}

func IsValidPassword(str string) bool {
	return regexp.MustCompile(passwordPattern).MatchString(str)
}
