package users

import (
	"net/http"
	"strconv"

	"superindo/database"
	"superindo/models"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	var deleteUser models.User

	userID := c.Param("userID")
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	deleteUser.ID = uint(id)

	err = database.DB.Delete(&deleteUser).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User's Account Has Been Deleted",
	})
}
