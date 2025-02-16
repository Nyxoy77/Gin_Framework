package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nyxoy77/gin-framework/authservice"
	"github.com/nyxoy77/gin-framework/database"
	"github.com/nyxoy77/gin-framework/migrations"
)

// type User struct {
// 	Name string `json:"name" binding:"required"`
// 	Age  int    `json:"age" `
// }

// var allUser = []User{
// 	{Name: "Shivam", Age: 21},
// }

func main() {
	fmt.Println("Learning GIN ")
	database.InitDb()
	migrations.RunMigrations()
	r := gin.Default()
	regRoutes := r.Group("/user")
	{
		regRoutes.POST("/register", authservice.RegisterUser)

	}
	r.Run(":8080")
}
