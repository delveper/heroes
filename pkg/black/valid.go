package black

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const defaultKey = "regex"

var ErrUnexpected = errors.New("unexpected error occurred")
var ErrValidating = errors.New("validation error")

// ErrorValidation is what it is
// we can catch it type in logistics level
type ErrorValidation struct {
	entity   string
	property string
	isZero   bool
	code     string
}

func (errV *ErrorValidation) Error() string {
	if !errV.isZero {
		errV.code = "valid" + " "
	}

	return strings.ToLower(fmt.Sprintf("%s has to have %s%s",
		errV.entity,
		errV.code,
		errV.property))
}

// ValidateStruct validates struct fields
// according to given regex tag
func ValidateStruct(src any) (err error) {
	// check if src is a struct
	srcValue, err := inspectSource(src)
	if err != nil {
		return err
	}
	// it is not abs necessary, just in case
	defer func() {
		if recover() != nil {
			err = ErrUnexpected
		}
	}()
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
					&ErrorValidation{entity: structName, property: fieldName, isZero: true})
			}
			// field validation according pattern
			if !regexp.MustCompile(pattern).MatchString(fmt.Sprintf("%v", fieldValue)) {
				return fmt.Errorf("%s: %w", ErrValidating,
					&ErrorValidation{entity: structName, property: fieldName})
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
	// in case of panic we will return ErrUnexpected
	return err
}
