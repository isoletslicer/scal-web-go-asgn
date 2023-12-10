package controllers

import (
	"assignment2/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InsertOrders(c *gin.Context) {
	var newOrder models.Orders
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	insertedOrder, err := InsertMethod(newOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data":    insertedOrder,
		"message": "Data successfully created",
		"success": true,
	})
}

func ShowOrders(c *gin.Context) {
	orders := GetAllOrdersMethod()
	c.JSON(http.StatusOK, gin.H{
		"Data": orders,
	})
}

func DeleteOrder(c *gin.Context) {
	OrderID := c.Param("OrderID")
	convertOrderID, err := strconv.Atoi(OrderID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":   nil,
			"status": fmt.Sprintf("%d", http.StatusNotFound),
		})
		return
	}
	DeleteIDMethod(uint(convertOrderID))
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Delete Data Order with ID %v success", OrderID),
		"success": true,
	})
}

func UpdateOrder(c *gin.Context) {
	var updatedOrder models.Orders
	OrderID := c.Param("OrderID")
	convertOrderID, err := strconv.Atoi(OrderID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":   nil,
			"status": fmt.Sprintf("%d", http.StatusNotFound),
		})
		return
	}
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	updatedOrder = UpdateIDMethod(updatedOrder, uint(convertOrderID))
	c.JSON(http.StatusOK, gin.H{
		"data":    updatedOrder,
		"message": fmt.Sprintf("Update Data Order with ID %v success", OrderID),
		"success": true,
	})

}
