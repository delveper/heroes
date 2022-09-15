// DISCLAIMER:
// all code below is written because of self-educational reasons
// and can be considered as huge overkill

package valid

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

const defaultKey = "regex"

var ErrInternal = errors.New("reflect is fun")

// StructRegex validates struct fields
// according to given regex tag
func StructRegex(src any) (err error) {
	valOf := reflect.Indirect(reflect.ValueOf(src))
	// struct structName
	var structName string
	if structName == "" {
		structName = valOf.Type().Name()
	}

	defer func() { // we do not panic but this part made just in case
		if r := recover(); r != nil {
			err = ErrInternal
		}
	}()

	for i := 0; i < valOf.NumField(); i++ {
		field := valOf.Field(i)
		// check re tag if any is present
		tags := valOf.Type().Field(i).Tag
		if tag, ok := getTag(tags, "regex"); ok {
			fieldVal := fmt.Sprintf("%v", field.Interface())
			if !regexp.MustCompile(tag).MatchString(fieldVal) {
				fieldName := valOf.Type().Field(i).Name
				return fmt.Errorf("%s has to have valid %s", structName, fieldName)
			}
		}
		// recursive call for nested structs
		if field.Type().Kind() == reflect.Struct {
			if err := StructRegex(field.Interface()); err != nil {
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
	pattern := fmt.Sprintf(`(?s)\s*(?P<key>%s):\"(?P<value>[^\"]+)\"`, defaultKey)

	re := regexp.MustCompile(pattern)
	if !re.MatchString(str) {
		return "", false
	}

	match := re.FindStringSubmatch(str)

	return match[len(match)-1], true
}
