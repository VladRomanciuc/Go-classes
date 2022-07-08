package router

import (
	"net/http"

	"github.com/VladRomanciuc/Go-classes/api/models"
	"github.com/go-chi/chi"

)

var chiRouter = chi.NewRouter()

type routerChi struct{}

func NewRouterChi() models.Router{
	return &routerChi{}
}

func (*routerChi) GET(url string, f func(w http.ResponseWriter, r *http.Request)){
	chiRouter.Get(url, f)
}
func (*routerChi)	POST(url string, f func(w http.ResponseWriter, r *http.Request)){
	chiRouter.Post(url, f)
}
func (*routerChi)	SERVE(port string){
	http.ListenAndServe(port, chiRouter)
}