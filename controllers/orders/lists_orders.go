package orders

import (
	"net/http"

	"superindo/database"
	"superindo/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ListsOrders(c *gin.Context) {
	var orders []models.Order

	userData := c.MustGet("userData").(jwt.MapClaims)
	id := uint(userData["id"].(float64))

	err := database.DB.Preload("Carts").Where("user_id = ?", id).Find(&orders).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "orders lists",
		"result":  orders,
	})
}
