package items

import (
	"net/http"

	"superindo/database"
	"superindo/helpers"
	"superindo/models"

	"github.com/gin-gonic/gin"
)

func CreateItem(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	Item := models.Item{}

	if contentType == helpers.AppJSON {
		c.ShouldBindJSON(&Item)
	} else {
		c.ShouldBind(&Item)
	}

	err := db.Debug().Create(&Item).Error
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Item Created",
	})
}
