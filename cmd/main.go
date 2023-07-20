package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"kiku-backend/database"
	"kiku-backend/routes"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "kiku"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")
	database.InitializeDatabase(db)
	routes.InitializeRoutes(db)
}
