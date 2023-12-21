package router

import (
	"finalprojekmygram/apps/controller"
	middleware "finalprojekmygram/apps/middlewares"
	"finalprojekmygram/apps/repository"
	"finalprojekmygram/apps/service"

	app "finalprojekmygram/external/database"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	db := app.ConnectDB()

	repo := repository.NewUserRepository(db)
	srv := service.NewUserService(repo)
	ctrl := controller.NewUserController(srv)

	userRouter := router.Group("/users")

	{
		userRouter.POST("/register", ctrl.Register)
		userRouter.POST("/login", ctrl.Login)

		{
			userRouter.PUT("/", middleware.Authentication(), ctrl.Update)
			userRouter.DELETE("/", middleware.Authentication(), ctrl.Delete)
		}
	}
}
