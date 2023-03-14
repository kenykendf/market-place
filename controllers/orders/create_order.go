package orders

import (
	"fmt"
	"net/http"

	"superindo/database"
	"superindo/helpers"
	"superindo/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	Order := models.Order{}

	userData := c.MustGet("userData").(jwt.MapClaims)
	id := uint(userData["id"].(float64))
	Order.UserID = id

	if contentType == helpers.AppJSON {
		c.ShouldBindJSON(&Order)
	} else {
		c.ShouldBind(&Order)
	}
	fmt.Println("ORDER ", Order)
	err := db.Debug().Create(&Order).Error
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Order Created",
	})
}
