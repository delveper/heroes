package black

import (
	"fmt"
	"reflect"
	"regexp"
)

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

func getStructName(src any) (string, error) {
	if err := inspectSource(src); err != nil {
		return "", err
	}
	return reflect.Indirect(reflect.ValueOf(src)).Type().Name(), nil
}

func getStructFieldValues(src any) (map[string]interface{}, error) {
	return nil, nil
}
