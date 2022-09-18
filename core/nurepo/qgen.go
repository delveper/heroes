package nurepo

import (
	"fmt"
	"os"
	"text/template"

	"github.com/delveper/heroes/pkg/black"
)

func genInsertQuery(src any) ([]byte, error) {
	const insert = `INSERT INTO "{{.Name}}" ( {{range .Data}}   {{.Tag}},   {{end}} )
	                             VALUES   ( {{range .Data}} '{{.Value}}', {{end}} )`

	var (
		tmpl *template.Template
		res  []black.StructValue
		err  error
	)

	if tmpl, err = template.New("fields").Parse(insert); err != nil {
		return nil, fmt.Errorf("error creating query template: %w", err)
	}

	if res, err = black.GetStructFieldValues(src); err != nil {
		return nil, fmt.Errorf("error parsing struct data: %w", err)
	}

	// prepare data
	data := struct {
		Name string
		Data []black.StructValue
	}{Name: "user", Data: res}

	if err = tmpl.Execute(os.Stdout, data); err != nil {
		return nil, fmt.Errorf("error composing query template: %w", err)
	}

	return nil, nil
}
