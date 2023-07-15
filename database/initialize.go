package database

import (
	"database/sql"
	"kiku-backend/database/migrations"
	"log"
)

func InitializeDatabase(db *sql.DB) error {
	err := migrations.InitializeUsers(db)
	if err != nil {
		log.Println("Failed to initialize users:", err)
		return err
	}
	err = migrations.InitializeMusic(db)
	if err != nil {
		log.Println("Failed to initialize music:", err)
		return err
	}
	return nil
}
