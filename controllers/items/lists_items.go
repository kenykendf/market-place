package items

import (
	"net/http"

	"superindo/database"
	"superindo/models"

	"github.com/gin-gonic/gin"
)

func ListsItems(c *gin.Context) {
	var items []models.Item

	query := c.Query("item_name")
	switch query {
	case "":
		err := database.DB.Find(&items).Error
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	default:
		err := database.DB.Where("item_name LIKE ?", "%"+query+"%").Find(&items).Error
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "items lists",
		"result":  items,
	})
}
