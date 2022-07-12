package main

import (
	"github.com/VladRomanciuc/Go-classes/api/service"
	"github.com/VladRomanciuc/Go-classes/api/dbapi"
	"github.com/VladRomanciuc/Go-classes/api/models"
	"github.com/VladRomanciuc/Go-classes/api/router"
	"github.com/VladRomanciuc/Go-classes/api/controller"
	"github.com/VladRomanciuc/Go-classes/api/cache"
)

var (
	//dbops models.DbOps = dbapi.NewFirestoreOps()
	dbops models.DbOps = dbapi.NewSQLiteDb()

	postService models.PostService = service.NewPostService(dbops)
	carDetailsService models.CarDetailsService = service.NewCarDetailsService()

	postCache models.PostCache = cache.NewRedisCache("localhost:49154", "redispw", 0, 60)

	postController models.PostController = controller.NewPostController(postService, postCache)
	carDetailsController models.CarDetailsController = controller.NewCarDetailsController(carDetailsService)
	
	//api models.Router = router.NewRouterMux()
	api models.Router = router.NewRouterChi()
)

func main() {
    port := ":8080"
	api.GET("/posts", postController.GetAll)
	api.POST("/posts", postController.AddPost)
	api.GET("/posts/{id}", postController.GetById)
	api.DELETE("/posts/{id}", postController.DeleteById)
    api.GET("/cardetails", carDetailsController.GetCarDetails)
	api.SERVE(port)
}