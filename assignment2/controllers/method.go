package controllers

import (
	"assignment2/database"
	"assignment2/models"
	"fmt"
)

func InsertMethod(Input models.Orders) (models.Orders, error) {
	db := database.GetDB()
	newOrder := Input
	db_err := db.Debug().Create(&newOrder).Error
	if db_err != nil {
		return models.Orders{}, db_err
	}
	fmt.Println("Data Added!")
	return newOrder, nil
}

func GetAllOrdersMethod() []models.Orders {
	db := database.GetDB()
	var orders []models.Orders
	db_err := db.Preload("Items").Find(&orders).Error
	if db_err != nil {
		panic(db_err)
	}
	return orders
}

func DeleteIDMethod(id uint) {
	db := database.GetDB()
	db_err := db.Where("order_id = ?", id).Delete(&models.Items{}).Error
	if db_err != nil {
		panic(db_err)
	}
	db_err = db.Delete(&models.Orders{}, id).Error
	if db_err != nil {
		panic(db_err)
	}
	fmt.Println("Data Deleted!")
}

func UpdateIDMethod(Input models.Orders, id uint) models.Orders {
	db := database.GetDB()
	updatedOrder := Input
	var err error
	for i := range updatedOrder.Items {
		err = db.Model(&updatedOrder.Items[i]).Where("item_id = ?", updatedOrder.Items[i].Item_ID).Updates(&updatedOrder.Items[i]).Error
		if err != nil {
			panic(err)
		}
	}
	var updated_Order models.Orders
	updated_Order.CustomerName = updatedOrder.CustomerName
	updated_Order.OrderedAt = updatedOrder.OrderedAt
	err = db.Model(&updated_Order).Where("order_id=?", id).Updates(&updated_Order).Error
	if err != nil {
		panic(err)
	}
	err = db.Preload("Items").Where("order_id=?", id).Find(&updatedOrder).Error
	if err != nil {
		panic(err)
	}
	fmt.Println("Data Updated!")
	return updatedOrder
}
