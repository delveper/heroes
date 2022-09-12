package ent

import (
	"regexp"
)

const (
	fullNamePattern = `^[\p{L}a-zA-Z&\s-'’.]{2,256}$`
	emailPattern    = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	passwordPattern = "^.{8,256}$"
)

func IsValidEmail(str string) bool {
	return regexp.MustCompile(emailPattern).MatchString(str)
}
func IsValidName(str string) bool {
	return regexp.MustCompile(fullNamePattern).MatchString(str)
}
func IsValidPassword(str string) bool {
	return regexp.MustCompile(passwordPattern).MatchString(str)
}
