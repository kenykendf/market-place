package categories

import (
	"net/http"
	"superindo/database"
	"superindo/models"

	"github.com/gin-gonic/gin"
)

func ListCategories(c *gin.Context) {
	var categories []models.Categories

	query := c.Query("categories")
	switch query {
	case "":
		err := database.DB.Preload("Item").Find(&categories).Error
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	default:
		err := database.DB.Preload("Item").Where("categories_name LIKE ?", "%"+query+"%").Find(&categories).Error
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Categories",
		"data":    categories,
	})
}
