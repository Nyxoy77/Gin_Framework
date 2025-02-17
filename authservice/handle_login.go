package authservice

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nyxoy77/gin-framework/database"
	"github.com/nyxoy77/gin-framework/models"
)

func HandleLogin(c *gin.Context) {
	var login models.Login

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	query := "SELECT password from Users where email = $1"
	var pass string
	if err := database.DB.QueryRow(context.Background(), query, login.Email).Scan(&pass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Either the username of password  is incorrect",
		})
		return
	}

	if login.Password == pass {
		c.JSON(http.StatusOK, gin.H{
			"message": "Loggged in successfully",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Either the username or password is incorrect",
		})
		return
	}
}
