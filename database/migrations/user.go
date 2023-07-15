package migrations

import (
	"database/sql"
	"fmt"
)

func InitializeUsers(db *sql.DB) error {
	_, err := db.Exec(`
    	CREATE TABLE IF NOT EXISTS users (
        	id SERIAL PRIMARY KEY,
        	login VARCHAR(50) UNIQUE NOT NULL,
        	password VARCHAR(255) NOT NULL,
        	email VARCHAR(255) UNIQUE NOT NULL
    	);
	`)
	if err != nil {
		return fmt.Errorf("Я не смог создать таблицу: %w", err)
	}
	return nil
}
