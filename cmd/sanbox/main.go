package main

import (
	"fmt"

	"github.com/delveper/heroes/pkg/black"
)

type Name struct {
	Password string `regex:"^.{8,255}$""`
}

type User struct {
	FullName string `regex:"^.{8,234}$"`
	Name
}

func main() {
	usr := false
	if err := black.ValidateStruct(usr); err != nil {
		fmt.Println(err)
	}

}
