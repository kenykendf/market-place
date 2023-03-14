package categories

import (
	"net/http"

	"superindo/database"
	"superindo/helpers"
	"superindo/models"

	"github.com/gin-gonic/gin"
)

func CreateCategories(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	Categories := models.Categories{}

	if contentType == helpers.AppJSON {
		c.ShouldBindJSON(&Categories)
	} else {
		c.ShouldBind(&Categories)
	}

	err := db.Debug().Create(&Categories).Error
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Category Created",
	})
}
