package controller

import (
	"github.com/gorilla/mux"
	"github.com/VladRomanciuc/Go-classes/api/demoapi"
)

//register func for the main paths and subrouter
func Register() *mux.Router {
    r := mux.NewRouter()
    api := r.PathPrefix("/api/v1").Subrouter()
	//api := r.PathPrefix("/api/v1").HeadersRegexp("Authorization", "Basic [a-zA-Z0-9]{1,128}").Subrouter()
	api.HandleFunc("/posts", demoapi.GetPosts).Methods("GET")
	api.HandleFunc("/posts", demoapi.AddPost).Methods("POST")
    return api
}

