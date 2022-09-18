package main

import (
	"fmt"

	"github.com/delveper/heroes/core/ent"
	"github.com/delveper/heroes/core/nurepo"
)

func main() {
	fmt.Println(nurepo.GenInsertQuery(ent.User{ID: "dffa", FullName: "basldfjffjasdfasf"}))
}
