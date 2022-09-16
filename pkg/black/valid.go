package black

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const defaultKey = "regex"

var (
	ErrUnexpected = errors.New("unexpected error has occurred")
)

// ValidationError is what it is
// we can catch it type in logistics level using errors.As()
type ValidationError struct {
	entity   string
	property string
}

func (err *ValidationError) Error() string {
	return strings.ToLower(fmt.Sprintf("%s has to have valid %v", err.entity, err.property))
}

// ValidateStruct validates struct fields
// according to given regex tag
func ValidateStruct(src any) (err error) {
	// not sure if we need this,
	// just in case something go wrong
	defer func() {
		if r := recover(); r != nil {
			err = ErrUnexpected
		}
	}()

	srcValue := reflect.Indirect(reflect.ValueOf(src))

	// only structs allowed
	var structName string
	if srcType := srcValue.Kind(); srcType != reflect.Struct {
		return fmt.Errorf("input value must be struct, got: %v", srcType)
	}

	// name of the top struct (we will scan all nested structs recursively)
	if structName == "" {
		structName = srcValue.Type().Name()
	}

	for i := 0; i < srcValue.NumField(); i++ {

		fieldValue := srcValue.Field(i)
		fieldType := srcValue.Type().Field(i)

		// check regex <key> if any is exists and retrieve its <value> => match[2]
		tagAll := fmt.Sprintf("%v", srcValue.Type().Field(i).Tag)
		tagValue := fmt.Sprintf(`(?s)(?i)\s*(?P<key>%s):\"(?P<value>[^\"]+)\"`, defaultKey)

		if match := regexp.MustCompile(tagValue).FindStringSubmatch(tagAll); match != nil {

			fieldToCheck := fmt.Sprintf("%v", fieldValue.Interface())
			pattern := match[2]

			// validate field value using given regex pattern
			if !regexp.MustCompile(pattern).MatchString(fieldToCheck) {
				return fmt.Errorf("error validating: %w",
					&ValidationError{entity: structName, property: fieldType.Name})
			}
		}

		// recursive call
		if fieldValue.Type().Kind() != reflect.Struct {
			continue
		}
		if err := ValidateStruct(fieldValue.Interface()); err != nil {
			return err
		}
	}

	return err
}
