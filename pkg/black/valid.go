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
var ErrValidating = errors.New("validation error")

// ValidationError is what it is
// we can catch it type in logistics level using errors.As()
// btw it is my first custom error
type ValidationError struct {
	entity   string
	property string
	isZero   bool
	code     string
}

func (err *ValidationError) Error() string {
	if !err.isZero {
		err.code = "valid "
		// TODO: Extend specific
	}

	return strings.ToLower(fmt.Sprintf("%s has to have %s%s",
		err.entity,
		err.code,
		err.property))
}

// ValidateStruct validates struct fields
// according to given regex tag
func ValidateStruct(src any) (err error) {
	if err := inspectSource(src); err != nil {
		return err
	}

	// this piece is not abs necessary
	defer func() {
		if recover() != nil {
			err = ErrUnexpected
		}
	}()

	srcValue := reflect.Indirect(reflect.ValueOf(src))

	// top level struct name (in case we are using nested structs)
	var structName string
	if structName == "" {
		structName = srcValue.Type().Name()
	}

	// iterate  all over struct fields
	for i := 0; i < srcValue.NumField(); i++ {

		fieldValue := srcValue.Field(i)
		fieldName := srcValue.Type().Field(i).Name
		tagValue := srcValue.Type().Field(i).Tag

		// check presence of regex tag (.Tag.Lookup() would not work here)
		if pattern, ok := GetTagValue(tagValue, defaultKey); ok {

			if fieldValue.IsZero() {
				return fmt.Errorf("%s: %w", ErrValidating,
					&ValidationError{entity: structName, property: fieldName, isZero: true})
			}

			// field validation according pattern
			if !regexp.MustCompile(pattern).MatchString(fmt.Sprintf("%v", fieldValue)) {
				return fmt.Errorf("%s: %w", ErrValidating,
					&ValidationError{entity: structName, property: fieldName})
			}
		}

		// recursive call for nested structs
		if fieldValue.Type().Kind() != reflect.Struct {
			continue
		}
		if err := ValidateStruct(fieldValue.Interface()); err != nil {
			return err
		}

	}

	return err
}
