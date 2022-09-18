package black

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const sqlKey = "sql"

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

type StructData struct {
	Name   string
	Fields []string
	Tags   []string
	Values []string
}

// GetStructData retrieve non-zero fields
// from structs with corresponding values that have `sql`tags
func GetStructData(src any) (*StructData, error) {
	if err := inspectSource(src); err != nil {
		return nil, err
	}

	srcValue := reflect.Indirect(reflect.ValueOf(src))

	res := new(StructData)

	if res.Name == "" {
		res.Name = strings.ToLower(srcValue.Type().Name())
	}

	for i := 0; i < srcValue.NumField(); i++ {

		if fieldValue := srcValue.Field(i); !fieldValue.IsZero() { // TODO: maybe it is redundant

			if fieldTag, ok := GetTagValue(srcValue.Type().Field(i).Tag, sqlKey); ok {
				res.Fields = append(res.Fields, srcValue.Type().Field(i).Name)
				res.Tags = append(res.Tags, fieldTag)
				res.Values = append(res.Values, fmt.Sprintf("%v", fieldValue.Interface()))

				// recursive call is allowed only for nested structs
				if fieldValue.Type().Kind() == reflect.Struct {
					add, err := GetStructData(fieldValue.Interface())
					if err != nil {
						return nil, err
					}

					res.Fields = append(res.Fields, add.Fields...)
					res.Tags = append(res.Tags, add.Tags...)
					res.Values = append(res.Values, add.Values...)
				}

			}

		}

	}

	return res, nil
}
