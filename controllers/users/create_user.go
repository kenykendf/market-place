package users

import (
	"net/http"

	"superindo/database"
	"superindo/helpers"
	"superindo/models"

	"github.com/gin-gonic/gin"
)

func CreatorUser(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "User Created",
		"username": User.Username,
		"email":    User.Email,
		"age":      User.Age,
	})
}
