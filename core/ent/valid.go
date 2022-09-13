package ent

import (
	"regexp"
)

const (
	fullNamePattern = `^[\p{L}a-zA-Z&\s-'’.]{2,256}$`
	emailPattern    = `^[A-Za-z0-9\.-]+@[A-Za-z]+\.[a-z]{2,3}$`
	passwordPattern = `^.{8,256}$` // do we need to put any constraint on passwords except length?
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
