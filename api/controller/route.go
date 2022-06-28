package controller

import (
	"/api/views"
	"encoding/json"
	"net/http"
)

func get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
    if r.Method == http.MethodGet {
		data := views.Response{
		Code: http.StatusOK,
		Body: "",
		}
		json.NewEncoder(w).Encode(data)
	}
    return
}

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
    return
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
    return
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
    return
}


func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("", get)
    mux.HandleFunc("", post)
    mux.HandleFunc("", put)
    mux.HandleFunc("", delete)
	return mux
}
