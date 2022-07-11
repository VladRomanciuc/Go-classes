package main

import (
	"github.com/VladRomanciuc/Go-classes/api/service"
	"github.com/VladRomanciuc/Go-classes/api/dbapi"
	"github.com/VladRomanciuc/Go-classes/api/models"
	"github.com/VladRomanciuc/Go-classes/api/router"
	"github.com/VladRomanciuc/Go-classes/api/controller"
)

var (
	//dbops models.DbOps = dbapi.NewFirestoreOps()
	dbops models.DbOps = dbapi.NewSQLiteDb()

	postService models.PostService = service.NewPostService(dbops)
	carDetailsService models.CarDetailsService = service.NewCarDetailsService()

	postController models.PostController = controller.NewPostController(postService)
	carDetailsController models.CarDetailsController = controller.NewCarDetailsController(carDetailsService)
	
	//api models.Router = router.NewRouterMux()
	api models.Router = router.NewRouterChi()
)

func main() {
    port := ":8080"
	api.GET("/posts", postController.GetAll)
	api.POST("/posts", postController.AddPost)
    api.GET("/cardetails", carDetailsController.GetCarDetails)
	api.SERVE(port)
}