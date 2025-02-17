package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nyxoy77/gin-framework/authservice"
	"github.com/nyxoy77/gin-framework/database"
	"github.com/nyxoy77/gin-framework/migrations"
)

func main() {
	err := godotenv.Load("cmd/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Learning GIN ")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	redirectURL := os.Getenv("REDIRECT_URL")

	// Print them to make sure they are correct
	fmt.Println("Client ID:", clientID)
	fmt.Println("Client Secret:", clientSecret)
	fmt.Println("Redirect URL:", redirectURL)

	// Check for empty values
	if clientID == "" || clientSecret == "" || redirectURL == "" {
		log.Fatal("One or more environment variables are missing or empty")
	}
	database.InitDb()
	migrations.RunMigrations()
	r := gin.Default()
	regRoutes := r.Group("/user")
	{
		regRoutes.POST("/register", authservice.RegisterUser)
		regRoutes.POST("/login", authservice.HandleLogin)
		regRoutes.GET("/oauth/login", authservice.HandleGoogleLogin)
		regRoutes.GET("/oauth/callback", authservice.HandleGoogleCallBack)

	}
	// r.Run(":8080")
	//Implementing the gracefull shutdown
	// server := &http.Server{
	// 	Addr:    ":8081",
	// 	Handler: r,
	// }

	// go func() {
	// 	fmt.Println("Server Started on the Port 8081")
	// 	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("An error occured listening the server %v", err)
	// 	}
	// }()
	// signalChan := make(chan os.Signal, 1)
	// signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	// <-signalChan // block the execution until the os signal occurs
	// fmt.Println("Gracefull Shutdown Begins...")
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := server.Shutdown(ctx); err != nil {
	// 	log.Fatal("Server Forced to shutdown with timeout")
	// }
	// fmt.Println("Server shut down gracefully")

	r.RunTLS(":8081", "server.crt", "server.key")
}
