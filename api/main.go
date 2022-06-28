package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
    "/api/controller"
    "github.com/gorilla/mux"
)

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

func main() {
    mux := controller.Register()
    r := mux.NewRouter()
    api := r.PathPrefix("/api/v1").Subrouter()

    api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)

    log.Fatal(http.ListenAndServe(":8080", api))
}