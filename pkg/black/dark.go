package black

import (
	"fmt"
	"reflect"
	"strings"
)

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
