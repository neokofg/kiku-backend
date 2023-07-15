package migrations

import (
	"database/sql"
	"fmt"
)

func InitializeMusic(db *sql.DB) error {
	_, err := db.Exec(`
    	CREATE TABLE IF NOT EXISTS music (
        	id SERIAL PRIMARY KEY,
        	name VARCHAR(50) NOT NULL,
        	author VARCHAR(255) NOT NULL,
        	url VARCHAR(255) UNIQUE NOT NULL,
    	    user_id INT NOT NULL,
    	    FOREIGN KEY (user_id) REFERENCES users (id)
    	);
	`)
	if err != nil {
		return fmt.Errorf("Я не смог создать таблицу: %w", err)
	}
	return nil
}
