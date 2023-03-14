package items

import (
	"net/http"
	"strconv"

	"superindo/database"
	"superindo/models"

	"github.com/gin-gonic/gin"
)

func DetailItem(c *gin.Context) {
	var item models.Item

	itemID := c.Param("itemID")
	id, err := strconv.ParseUint(itemID, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	item.ID = uint(id)

	err = database.DB.First(&item, item.ID).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "item detail",
		"result":  item,
	})
}
