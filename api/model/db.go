package model

import (
	"fmt"
	"log"

	"github.com/VladRomanciuc/Go-classes/api/views"
	"github.com/objectbox/objectbox-go/objectbox"
)
var con *objectbox.ObjectBox

func Connect() *objectbox.ObjectBox {
		objectBox, err := objectbox.NewBuilder().Model(views.Task.ObjectBoxModel()).Build()
		if err !=nil {
			log.Fatal(err)
		}
		fmt.Println("Connected to DB")
		con = objectBox
		return con
}