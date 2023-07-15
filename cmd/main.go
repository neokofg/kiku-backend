package main

import (
	"database/sql"
	"fmt"
	"kiku-backend/app/application"
	"kiku-backend/app/handler"
	"kiku-backend/infrastructure"
	"log"
	"net/http"

	_ "github.com/lib/pq"
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
	repo := &infrastructure.SqlUserRepository{DB: db}
	service := &application.RegisterService{Repo: repo}
	loginService := &application.LoginService{Repo: repo}
	getUserHandler := &handler.GetUserHandler{LoginService: loginService}
	http.Handle("/register", handler.RegisterHandler(*service))
	http.Handle("/login", handler.LoginHandler(*loginService))
	http.Handle("/user", getUserHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
