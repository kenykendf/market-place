package users

import (
	"fmt"
	"net/http"
	"strconv"

	"superindo/database"
	"superindo/models"

	"github.com/gin-gonic/gin"
)

func ListsUsers(c *gin.Context) {
	var user []models.User

	limit := c.Query("limit")
	var limitpage int
	if limit != "" {
		limit = "10"
		limitpage, _ = strconv.Atoi(limit)
		fmt.Println("limitpage", limitpage)

	}

	offset := c.Query("offset")
	var offsetpage int
	if offset != "" {
		offset = "0"
		offsetpage, _ = strconv.Atoi(offset)
		fmt.Println("offsetpage", offsetpage)
	}

	emailQuery := c.Query("email")

	switch emailQuery {
	case "":
		err := database.DB.Offset(offsetpage).Find(&user).Limit(limitpage).Error
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	default:
		err := database.DB.Where("email LIKE ?", "%"+emailQuery+"%").Find(&user).Error
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user lists",
		"result":  user,
	})
}
