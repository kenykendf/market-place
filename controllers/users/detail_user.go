package users

import (
	"net/http"
	"strconv"

	"superindo/database"
	"superindo/models"

	"github.com/gin-gonic/gin"
)

func DetailUser(c *gin.Context) {
	var user models.User

	userID := c.Param("userID")
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user.ID = uint(id)

	err = database.DB.First(&user, user.ID).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "finder pong",
		"result":  user,
	})
}
