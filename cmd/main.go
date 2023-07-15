package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"kiku-backend/app/application"
	"kiku-backend/app/handler"
	"kiku-backend/infrastructure"
	"log"
	"net/http"
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

	// create an http.ServeMux
	mux := http.NewServeMux()

	// handle the routes
	mux.Handle("/register", handler.RegisterHandler(*service))
	mux.Handle("/user", getUserHandler)
	mux.Handle("/login", handler.LoginHandler(*loginService))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // your app's origin
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})

	corsHandler := c.Handler(mux)

	// start the server
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
