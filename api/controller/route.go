package controller

import (
	"encoding/json"
	"net/http"

	"github.com/VladRomanciuc/Go-classes/api/views"
	"github.com/gorilla/mux"
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
	api.HandleFunc("", get)
    api.HandleFunc("", post)
    api.HandleFunc("", put)
    api.HandleFunc("", delete)
	//mux.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)
    return api
}

/*
func params(w http.ResponseWriter, r *http.Request) {
   
    pathParams := mux.Vars(r)
    w.Header().Set("Content-Type", "application/json")

    userID := -1
    var err error
    if val, ok := pathParams["userID"]; ok {
        userID, err = strconv.Atoi(val)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(`{"message": "need a number"}`))
            return
        }
    }

    commentID := -1
    if val, ok := pathParams["commentID"]; ok {
        commentID, err = strconv.Atoi(val)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(`{"message": "need a number"}`))
            return
        }
    }

    query := r.URL.Query()
    location := query.Get("location")

    w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s" }`, userID, commentID, location)))
}
*/