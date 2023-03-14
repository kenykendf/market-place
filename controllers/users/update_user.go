package users

import (
	"net/http"
	"strconv"

	"superindo/database"
	"superindo/helpers"
	"superindo/models"

	"github.com/gin-gonic/gin"
)

func UpdaterUser(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	User := models.User{}

	userID, _ := strconv.Atoi(c.Param("userID"))

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	User.ID = uint(userID)

	err := db.Model(&User).Where("id = ?", userID).Updates(models.User{
		Username: User.Username,
		Email:    User.Email,
		Password: User.Password,
		Age:      User.Age,
		Fullname: User.Fullname,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, User)
}
