package main

import (
	"github.com/VladRomanciuc/Go-classes/api/router"
	"github.com/VladRomanciuc/Go-classes/api/controller"
)

var (
	postController controller.PostController = controller.NewPostController()
	routerMux router.Router = router.NewRouterMux()
)

func main() {
    port := ":8080"
	routerMux.GET("/posts", postController.GetAll)
	routerMux.POST("/posts", postController.AddPost)
    routerMux.SERVE(port)
}