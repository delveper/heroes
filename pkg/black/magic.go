package black

import (
	"fmt"
	"reflect"
	"regexp"
)

// GetTagValue is designed because luck of functionality in reflect.Tag.Lookup()
// and help retrieve <value> in given <key> from struct fields
func GetTagValue(tag reflect.StructTag, key string) (string, bool) {
	tagStr := fmt.Sprintf("%v", tag)
	tagValue := fmt.Sprintf(`(?s)(?i)\s*(?P<key>%s):\"(?P<value>[^\"]+)\"`, key)

	if match := regexp.MustCompile(tagValue).
		FindStringSubmatch(tagStr); match != nil {
		return match[2], true
	}
	return "", false
}

func inspectSource(src any) (err error) {
	defer func() {
		if recover() != nil {
			err = ErrUnexpected
		}
	}()

	srcValue := reflect.Indirect(reflect.ValueOf(src))

	if srcType := srcValue.Kind(); srcType != reflect.Struct {
		return fmt.Errorf("input value must be struct, got: %v", srcType)
	}

	return err
}

func GetStructName(src any) (string, error) {
	if err := inspectSource(src); err != nil {
		return "", err
	}
	return reflect.Indirect(reflect.ValueOf(src)).Type().Name(), nil
}

// StructValue will be used for code gen
type StructValue struct {
	Field string
	Tag   string
	Value interface{}
}

// GetStructFieldValues retrieve non-zero fields
// from structs with corresponding values that have `sql`tags
func GetStructFieldValues(src any) ([]StructValue, error) {
	if err := inspectSource(src); err != nil {
		return nil, err
	}

	srcValue := reflect.Indirect(reflect.ValueOf(src))

	res := make([]StructValue, 0, srcValue.NumField())

	for i := 0; i < srcValue.NumField(); i++ {
		var val StructValue

		if fieldValue := srcValue.Field(i); !fieldValue.IsZero() {

			if fieldTag, ok := GetTagValue(srcValue.Type().Field(i).Tag, "sql"); ok {

				val.Field = srcValue.Type().Field(i).Name
				val.Value = fieldValue.Interface()
				val.Tag = fieldTag

				res = append(res, val)

				// recursive call is allowed only for nested structs
				if fieldValue.Type().Kind() != reflect.Struct {
					continue
				}

				add, err := GetStructFieldValues(fieldValue.Interface())
				if err != nil {
					return nil, err
				}

				res = append(res, add...)
			}

		}

	}

	return res, nil
}

// TODO: Overkill :)
// func HasZeroValue[T ~int | ~string](val T) bool {
// 	return val == *new(T)
// }
