package main

import (
	"github.com/VladRomanciuc/Go-classes/api/service"
	"github.com/VladRomanciuc/Go-classes/api/dbapi"
	"github.com/VladRomanciuc/Go-classes/api/models"
	"github.com/VladRomanciuc/Go-classes/api/router"
	"github.com/VladRomanciuc/Go-classes/api/controller"
)

var (
	dbops models.DbOps = dbapi.NewPostOpsCollection()
	postService service.PostService = service.NewPostService(dbops)
	postController controller.PostController = controller.NewPostController(postService)
	//api models.Router = router.NewRouterMux()
	api models.Router = router.NewRouterChi()
)

func main() {
    port := ":8080"
	api.GET("/posts", postController.GetAll)
	api.POST("/posts", postController.AddPost)
    api.SERVE(port)
}