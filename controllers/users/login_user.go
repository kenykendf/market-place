package users

import (
	"net/http"

	"superindo/database"
	"superindo/helpers"
	"superindo/models"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func LoginUser(c *gin.Context) {
	db := database.GetDB()
	contenType := helpers.GetContentType(c)
	_, _ = db, contenType
	User := models.User{}
	password := ""

	if contenType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid Credential",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
