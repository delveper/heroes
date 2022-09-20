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

func inspectSource(src any) (val *reflect.Value, err error) {
	defer func() {
		if recover() != nil {
			err = ErrUnexpected
		}
	}()

	srcValue := reflect.Indirect(reflect.ValueOf(src))
	val = &srcValue

	if srcType := srcValue.Kind(); srcType != reflect.Struct {
		return nil, fmt.Errorf("input value must be struct, got: %v", srcType)
	}

	return val, err
}

// GetStructName retrieve name of underlying struct
func GetStructName(src any) (string, error) {
	srcValue, err := inspectSource(src)
	if err != nil {
		return "", err
	}
	return srcValue.Type().Name(), nil
}
