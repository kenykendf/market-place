package orders

import (
	"net/http"
	"strconv"

	"superindo/database"
	"superindo/models"

	"github.com/gin-gonic/gin"
)

func DetailOrder(c *gin.Context) {
	var order models.Order

	orderID := c.Param("orderID")
	id, err := strconv.ParseUint(orderID, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	order.ID = uint(id)

	err = database.DB.Preload("Items").First(&order, order.ID).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "order detail",
		"result":  order,
	})
}
