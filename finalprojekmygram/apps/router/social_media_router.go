package router

import (
	"finalprojekmygram/apps/controller"
	middleware "finalprojekmygram/apps/middlewares"
	"finalprojekmygram/apps/repository"
	"finalprojekmygram/apps/service"

	app "finalprojekmygram/external/database"

	"github.com/gin-gonic/gin"
)

func SocialMediaRouter(router *gin.Engine) {
	db := app.ConnectDB()

	repo := repository.NewSocialMediaRepository(db)
	srv := service.NewSocialMediaService(repo)
	ctrl := controller.NewSocialMediaController(srv)

	socialMedia := router.Group("/socialmedias", middleware.Authentication())

	{
		socialMedia.GET("/", ctrl.GetAll)
		socialMedia.POST("/", ctrl.Create)
		{
			socialMedia.PUT("/:id", middleware.SocialMediaAuthz(srv), ctrl.Update)
			socialMedia.DELETE("/:id", middleware.SocialMediaAuthz(srv), ctrl.Delete)
		}
	}
}
