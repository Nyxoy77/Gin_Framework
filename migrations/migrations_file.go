package migrations

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // ✅ Explicitly import the file driver
)

func RunMigrations() {
	dbURL := "postgres://postgres:hello@localhost:5432/postgres?sslmode=disable"

	// ✅ Use the correct migration path based on your folder structure
	migrationsPath := `file://C:/Users/Shivam%20Rai/OneDrive/Desktop/DESKTOP/go_lco/Project/gin/migrations/`

	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		log.Fatalf("Migration setup failed: %v\n", err)
		return
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v\n", err)
		return
	}

	log.Println("Migrations applied successfully!")
}
