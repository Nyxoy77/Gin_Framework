package authservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nyxoy77/gin-framework/database"
	"github.com/nyxoy77/gin-framework/models"
)

func RegisterUser(c *gin.Context) {

	time.Sleep(10 * time.Second)
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
	// To avoid sql injection we should not directly pass the query !
	// The username and the email must be unique!
	if err := database.DB.QueryRow(context.Background(), query, user.username, user.Email, user.Password).Scan(&userId); err != nil {
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
