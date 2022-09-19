package main

import (
	"fmt"

	"github.com/delveper/heroes/core/ent"
	"github.com/delveper/heroes/pkg/black"
)

func main() {
	usr := ent.User{
		ID:       "dfdfdfa",
		FullName: "sdfasd",
		Email:    "sadfas@cxvx.df",
		Password: "",
	}
	fmt.Println(black.ValidateStruct(usr))
}
