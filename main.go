package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nyxoy77/gin-framework/authservice"
	"github.com/nyxoy77/gin-framework/database"
	"github.com/nyxoy77/gin-framework/migrations"
)

func main() {
	fmt.Println("Learning GIN ")
	database.InitDb()
	migrations.RunMigrations()
	r := gin.Default()
	regRoutes := r.Group("/user")
	{
		regRoutes.POST("/register", authservice.RegisterUser)

	}
	// r.Run(":8080")
	//Implementing the gracefull shutdown
	server := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	go func() {
		fmt.Println("Server Started on the Port 8081")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("An error occured listening the server %v", err)
		}
	}()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan // block the execution until the os signal occurs
	fmt.Println("Gracefull Shutdown Begins...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Forced to shutdown with timeout")
	}
	fmt.Println("Server shut down gracefully")
}
