package router

import (
	"net/http"

	"github.com/gorilla/mux"

)

var router = mux.NewRouter()

type routerMux struct{}

func NewRouterMux() Router{
	return &routerMux{}
}

func (*routerMux) GET(url string, f func(w http.ResponseWriter, r *http.Request)){
	router.HandleFunc(url, f).Methods("GET")
}
func (*routerMux)	POST(url string, f func(w http.ResponseWriter, r *http.Request)){
	router.HandleFunc(url, f).Methods("POST")
}
func (*routerMux)	SERVE(port string){
	http.ListenAndServe(port, router)
}