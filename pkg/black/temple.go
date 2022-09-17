package black

import (
	"reflect"
)

type Entity interface {
	Name() string
	NonZeroFields() map[string]interface{}
}

func HasZeroValue(src any) bool {
	valOf := reflect.Indirect(reflect.ValueOf(src))
	return valOf.IsZero()
}

func HasZeroValueGen[T ~int | ~string](val T) bool {
	return val == *new(T)
}
