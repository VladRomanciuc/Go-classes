package controller

import (
	"encoding/json"
	"net/http"
	"github.com/VladRomanciuc/Go-classes/api/model"
	"github.com/VladRomanciuc/Go-classes/api/views"
	"github.com/gorilla/mux"
)


func post(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    if r.Method == http.MethodPost {
		data := views.Response{
		Code: http.StatusCreated,
		Body: "",
		}
		json.NewEncoder(w).Encode(data)
	}
}

func put(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusAccepted)
    if r.Method == http.MethodPost {
		data := views.Response{
		Code: http.StatusAccepted,
		Body: "",
		}
		json.NewEncoder(w).Encode(data)
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if r.Method == http.MethodPost {
		data := views.Response{
		Code: http.StatusOK,
		Body: "",
		}
		json.NewEncoder(w).Encode(data)
	}
}


func Register() *mux.Router {
    r := mux.NewRouter()
    api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", model.Get)
    api.HandleFunc("", post)
    api.HandleFunc("", put)
    api.HandleFunc("", delete)
    return api
}
