package ent

import (
	"regexp"
)

const (
	fullNamePattern = "^[\\p{L}a-zA-Z&\\s-'’.]{2,255}$"
	emailPattern    = "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	passwordPattern = "^.{8,255}$" // do we need to put any constraint on passwords except length?
	maxLength       = 255
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
