package authservice

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nyxoy77/gin-framework/database"
	"github.com/nyxoy77/gin-framework/models"
)

func RegisterUser(c *gin.Context) {

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(user)
	// Query to database
	query := `INSERT into users(username,email,password) VALUES ($1,$2,$3) RETURNING id`
	var userId int
	if err := database.DB.QueryRow(context.Background(), query, user.Name, user.Email, user.Password).Scan(&userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed to register the user",
			"err":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registerd the user!",
		"user_id": userId,
	})

}
