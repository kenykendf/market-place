package orders

import (
	"fmt"
	"net/http"

	"superindo/database"
	"superindo/models"

	"github.com/gin-gonic/gin"
)

func ListsOrders(c *gin.Context) {
	var orders []models.Order

	query := c.Query("email")
	switch query {
	case "":
		err := database.DB.Preload("Carts").Find(&orders).Error
		fmt.Println("ERR CHECK ", err)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	default:
		err := database.DB.Preload("Carts").Where("email LIKE ?", "%"+query+"%").Find(&orders).Error
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "orders lists",
		"result":  orders,
	})
}
