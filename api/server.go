package main

import (
	"github.com/VladRomanciuc/Go-classes/api/models"
	"github.com/VladRomanciuc/Go-classes/api/router"
	"github.com/VladRomanciuc/Go-classes/api/controller"
)

var (
	postController controller.PostController = controller.NewPostController()
	//api models.Router = router.NewRouterMux()
	api models.Router = router.NewRouterChi()
)

func main() {
    port := ":8080"
	api.GET("/posts", postController.GetAll)
	api.POST("/posts", postController.AddPost)
    api.SERVE(port)
}