package main

import (
	"fmt"
	"time"

	"github.com/delveper/heroes/core/ent"
	"github.com/delveper/heroes/pkg/black"
)

type User ent.User

func main() {
	usr := ent.User{
		FullName:  "dfdfafasdf",
		Email:     "df@dsf.com",
		Password:  "asd",
		CreatedAt: time.Time{},
	}

	fmt.Println(black.GetStructFieldValues(usr))

	// usr := User{
	// 	FullName: "No name",
	// 	Email:    "email",
	// }
	// tmp, err := template.New("add").Parse(`SELECT * FROM "{{.FullName}}" WHERE email = {{.Email}} {{.Name}}`)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// if err := tmp.Execute(os.Stdout, usr); err != nil {
	// 	log.Fatal(err)
	// }
}
