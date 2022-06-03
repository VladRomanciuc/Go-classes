package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//func to set header and type json
func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("{}"))
}


func main() {
	//set the handler as ServeMux
	m := http.NewServeMux()
	m.HandleFunc("/", testHandler)

	//set the port
	const addr = "localhost:8080"
	
	//define the server
	srv := http.Server{
		Handler:      m,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	//print the info before start the server
	fmt.Println("server started on ", addr)
	
	//start the server with error handler
	err := srv.ListenAndServe()
	log.Fatal(err)

}

