// DISCLAIMER:
// all code below is full of black magic
// and was done a clumsy way
// btw it works well and gives more flexibility
// in returning validation errors

package black

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const defaultKey = "regex"

var ErrUnexpected = errors.New("unexpected error has occurred")

// ValidateStruct validates struct fields
// according to given regex tag
func ValidateStruct(src any) (err error) {
	valOf := reflect.Indirect(reflect.ValueOf(src))

	var structName string
	if structName == "" {
		structName = strings.ToLower(valOf.Type().Name())
	}

	defer func() { // just in case
		if r := recover(); r != nil {
			err = ErrUnexpected
		}
	}()

	for i := 0; i < valOf.NumField(); i++ {
		field := valOf.Field(i)
		// check re tag if any is present
		tags := valOf.Type().Field(i).Tag
		if tag, ok := getTag(tags, "regex"); ok {
			fieldVal := fmt.Sprintf("%v", field.Interface())

			if !regexp.MustCompile(tag).MatchString(fieldVal) {
				fieldName := strings.ToLower(valOf.Type().Field(i).Name)
				return fmt.Errorf("%s has to have valid %s", structName, fieldName)
			}
		}
		// recursive call for nested structs
		if field.Type().Kind() == reflect.Struct {
			if err := ValidateStruct(field.Interface()); err != nil {
				return err
			}
		}
	}

	return err
}

// getTag help to improve reflect StructField Lookup() method
// that did not meet our expectations
func getTag(tag reflect.StructTag, key string) (string, bool) {
	str := fmt.Sprintf("%v", tag)
	pattern := fmt.Sprintf(`(?s)(?i)\s*(?P<key>%s):\"(?P<value>[^\"]+)\"`, defaultKey)

	re := regexp.MustCompile(pattern)
	if !re.MatchString(str) {
		return "", false
	}

	match := re.FindStringSubmatch(str)

	return match[2], true
}
