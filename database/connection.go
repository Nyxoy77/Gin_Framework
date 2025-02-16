package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDb() {
	dbURL := "postgres://postgres:hello@localhost:5432/postgres?sslmode=disable"

	var err error
	DB, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	fmt.Println("Connected to PostgreSQL successfully!")
}

func DbClose() {
	DB.Close()
	fmt.Println("Db closed")
}
