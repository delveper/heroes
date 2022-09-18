package nurepo

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/delveper/heroes/pkg/black"
)

func GenInsertQuery(src any) ([]byte, error) {
	const insert = `INSERT INTO "{{.Name}}" ( {{ StringsJoin .Tags ", "     }} )
	                               VALUES   ('{{ StringsJoin .Values "', '" }}');`
	var (
		tmpl *template.Template
		res  *black.StructData
		buf  bytes.Buffer
		err  error
	)

	if tmpl, err = template.New("insert").
		Funcs(template.FuncMap{"StringsJoin": strings.Join}).
		Parse(insert); err != nil {
		return nil, fmt.Errorf("error creating query template: %w", err)
	}

	if res, err = black.GetStructData(src); err != nil {
		return nil, fmt.Errorf("error parsing struct data: %w", err)
	}

	if err = tmpl.Execute(&buf, res); err != nil {
		return nil, fmt.Errorf("error composing query template: %w", err)
	}
	fmt.Println(string(buf.Bytes()))
	return buf.Bytes(), nil
}
