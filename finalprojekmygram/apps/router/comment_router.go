package router

import (
	"finalprojekmygram/apps/controller"
	middleware "finalprojekmygram/apps/middlewares"
	"finalprojekmygram/apps/repository"
	"finalprojekmygram/apps/service"

	app "finalprojekmygram/external/database"

	"github.com/gin-gonic/gin"
)

func CommentRouter(router *gin.Engine) {
	db := app.ConnectDB()

	repoPhoto := repository.NewPhotoRepository(db)
	srvPhoto := service.PhotoService(repoPhoto)

	repoComment := repository.NewCommentRepository(db)
	srvComment := service.NewCommentService(repoComment)

	ctrl := controller.NewCommentController(srvComment, srvPhoto)

	commentRouter := router.Group("/comments", middleware.Authentication())

	{
		commentRouter.POST("/", ctrl.Create)
		commentRouter.GET("/", ctrl.GetAll)
		{
			commentRouter.PUT("/:commentId", middleware.CommentAuthz(srvComment), ctrl.Update)
			commentRouter.DELETE("/:commentId", middleware.CommentAuthz(srvComment), ctrl.Delete)
		}
	}
}
