package controller

import (
	"github.com/gorilla/mux"
	"github.com/VladRomanciuc/Go-classes/api/firestoreapi"
)

//register func for the main paths and subrouter
func Register() *mux.Router {
    r := mux.NewRouter()
    api := r.PathPrefix("/api/v1").Subrouter()
	//api := r.PathPrefix("/api/v1").HeadersRegexp("Authorization", "Basic [a-zA-Z0-9]{1,128}").Subrouter()
	api.HandleFunc("/posts", firestoreapi.GetPosts).Methods("GET")
	api.HandleFunc("/posts", firestoreapi.AddPost).Methods("POST")
    return api
}

