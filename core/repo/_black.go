package repo

import (
	"fmt"
	"reflect"
	"strings"
)

// DISCLAIMER
// experimental
// huge overkill
// haven't used in project and probably never will

const tagSQL = "sql" // custom tag, that we define in ent.User fields

type Data struct {
	Name        string
	FieldNames  []string
	FieldValues []string // []interface{} ?
	TagNames    []string
}

// AddAgnostic can dynamically build SQL,
// TODO: Error handling
func (k *Keeper) AddAgnostic(src any) error {
	res := StructData(src, tagSQL)

	name := strings.ToLower(res.Name)          // table name
	fields := strings.Join(res.TagNames, `, `) // table fields ...
	values := fmt.Sprintf("'%s'", strings.Join(res.FieldValues, `', '`))

	SQL := fmt.Sprintf(`INSERT INTO "%s" (%s) VALUES(%s);`, name, fields, values)
	k.repo.QueryRow(SQL)

	return nil
}

// StructData gives struct data that we can use
// to dynamically work along with SQL
func StructData(src any, key string) *Data {
	res := &Data{}
	valueOf := reflect.Indirect(reflect.ValueOf(src))

	if res.Name == "" { // recursively we will get another struct name
		res.Name = valueOf.Type().Name()
	}

	for i := 0; i < valueOf.NumField(); i++ {
		field := valueOf.Field(i)

		res.FieldNames = append(res.FieldNames, valueOf.Type().Field(i).Name)
		res.FieldValues = append(res.FieldValues, fmt.Sprint(field.Interface()))
		res.TagNames = append(res.TagNames, valueOf.Type().Field(i).Tag.Get(key))

		// recursive call for nested structs
		if field.Type().Kind() == reflect.Struct && field.CanAddr() {
			StructData(field.Addr().Interface(), key)
		}
	}
	return res
}
