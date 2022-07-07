package main

import (
	"log"
	"net/http"
	"github.com/VladRomanciuc/Go-classes/api/controller"
)


func main() {
    api := controller.Register() 
    log.Fatal(http.ListenAndServe(":8080", api))
}