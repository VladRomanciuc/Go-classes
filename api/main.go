package main

import (
    "log"
    "net/http"
    "github.com/VladRomanciuc/Go-classes/api/controller"
    //"github.com/VladRomanciuc/Go-classes/api/model"
)


func main() {
    api := controller.Register()
    //db := model.Connect()
    //defer db.Close()
    log.Fatal(http.ListenAndServe(":8080", api))
}