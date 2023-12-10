package routers

import (
	ctrl "assignment2/controllers"

	"github.com/gin-gonic/gin"
)

func ServerInit() *gin.Engine {
	router := gin.Default()
	router.POST("/api/orders", ctrl.InsertOrders)
	router.GET("/api/orders", ctrl.ShowOrders)
	router.PUT("/api/orders/:OrderID", ctrl.UpdateOrder)
	router.DELETE("/api/orders/:OrderID", ctrl.DeleteOrder)
	return router
}
