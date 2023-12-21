package router

import (
	"finalprojekmygram/apps/controller"
	middleware "finalprojekmygram/apps/middlewares"
	"finalprojekmygram/apps/repository"
	"finalprojekmygram/apps/service"

	app "finalprojekmygram/external/database"

	"github.com/gin-gonic/gin"
)

func PhotoRouter(router *gin.Engine) {
	db := app.ConnectDB()

	repo := repository.NewPhotoRepository(db)
	srv := service.NewPhotoService(repo)
	ctrl := controller.NewPhotoController(srv)

	photoRouter := router.Group("/photos", middleware.Authentication())

	{
		photoRouter.POST("/", ctrl.Create)
		photoRouter.GET("/", ctrl.GetAll)
		{
			photoRouter.PUT("/:id", middleware.PhotoAuthz(srv), ctrl.Update)
			photoRouter.DELETE("/:id", middleware.PhotoAuthz(srv), ctrl.Delete)
		}
	}
}
