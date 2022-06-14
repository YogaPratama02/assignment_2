package controllers

import (
	"BootcampHacktiv8/assignment_2/db"
	"BootcampHacktiv8/assignment_2/models"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

func IndexController(c *gin.Context) {
	db := db.DbManager()

	orders := []models.Order{}
	db.Preload("Item").Find(&orders)
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "Successfully get orders",
		"data":    &orders,
	})
}

func CheckItem(value interface{}) error {
	data, ok := value.(models.Order)
	if !ok {
		return errors.New("Error validation")
	}
	for _, value := range data.Item {
		if value.ItemCode == "" {
			return errors.New("Error Item Code")
		}
	}
	return nil
}

func CreateController(c *gin.Context) {
	db := db.DbManager()
	// order := models.Order{}
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  false,
			"message": "Error read body",
			"data":    nil,
		})
		return
	}

	validate := validation.ValidateStruct(&order,
		validation.Field(&order.CustomerName, validation.Required),
		validation.Field(&order.OrderedAt, validation.Required),
		// validation.By(CheckItem),
	)
	errs := validation.Validate(order, validation.By(CheckItem))
	if errs != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  false,
			"message": errs.Error(),
			"data":    nil,
		})
		return
	}

	if validate != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  false,
			"message": validate,
			"data":    nil,
		})
		return
	}

	err := db.Create(&order)
	if err.Error != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  false,
			"message": err.Error,
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"status":  true,
		"message": "Successfully create orders",
		"data":    nil,
	})
}

func UpdateController(c *gin.Context) {
	db := db.DbManager()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  false,
			"message": "Error convert data",
		})
		return
	}

	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  false,
			"message": "Error read body",
			"data":    nil,
		})
		return
	}

	validate := validation.ValidateStruct(&order,
		validation.Field(&order.CustomerName, validation.Required),
		validation.Field(&order.OrderedAt, validation.Required),
	)

	if validate != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  false,
			"message": validate,
			"data":    nil,
		})
		return
	}

	db.Model(&order).Where("id = ?", id).Updates(order)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, map[string]interface{}{
	// 		"status":  false,
	// 		"message": err,
	// 		"data":    nil,
	// 	})
	// 	return
	// }

	// err = db.Model(&order).Where("id = ?", id).Updates(order).Error
	for _, result := range order.Item {
		fmt.Println("Items", result)
		var item models.Item
		err = db.Model(&item).Where("id = ?", result.Id).Updates(result).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status":  false,
				"message": err,
				"data":    nil,
			})
			return
		}
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  false,
		"message": "Sucessfully update orders",
		"data":    nil,
	})
}

func DeleteController(c *gin.Context) {
	db := db.DbManager()

	var order models.Order

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  false,
			"message": "Error convert data",
		})
		return
	}

	checkId := db.Where("id = ?", id).First(&order)
	if checkId.Error != nil {
		log.Printf("Error get data with err: %s", checkId.Error)
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"data":    nil,
			"status":  false,
			"message": "Id Not Found",
		})
		return
	}

	db.Delete(order, id)

	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  false,
		"message": "Sucessfully delete orders",
		"data":    nil,
	})
}
