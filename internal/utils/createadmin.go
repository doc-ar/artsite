package utils

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateAdmin(dbpool *pgxpool.Pool, ctx context.Context) {
	row := dbpool.QueryRow(ctx, "SELECT COUNT(*) FROM admins")
	var count int
	err := row.Scan(&count)
	if err != nil {
		log.Fatal("Failed to check admin existence:", err)
	}

	if count == 0 {
		password := os.Getenv("ADMIN_PASSWORD") // Or use a secret manager
		if password == "" {
			log.Fatal("ADMIN_PASSWORD env variable not set")
		}

		hash, err := HashPassword(password)
		if err != nil {
			log.Fatal("Failed to hash admin password:", err)
		}

		_, err = dbpool.Exec(ctx, `
		INSERT INTO admins (username, password_hash)
		VALUES ($1, $2)
	`, "admin", hash)
		if err != nil {
			log.Fatal("Failed to insert admin user:", err)
		}

		log.Println("Default admin user created.")
	} else {
		log.Println("Admin user already exists. Skipping creation.")
	}
}
