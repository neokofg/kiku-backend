package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"kiku-backend/app/application"
	"kiku-backend/app/handler"
	"kiku-backend/database"
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
	database.InitializeDatabase(db)
	userRepo := &infrastructure.SqlUserRepository{DB: db}
	musicRepo := &infrastructure.SqlMusicRepository{DB: db}
	registerService := &application.RegisterService{Repo: userRepo}
	loginService := &application.LoginService{Repo: userRepo}
	musicService := &application.MusicService{Repo: musicRepo, UserRepository: userRepo}
	getUserHandler := &handler.GetUserHandler{LoginService: loginService}
	getFileHandler := &handler.GetFileHandler{}

	// create an http.ServeMux
	mux := http.NewServeMux()

	// handle the routes
	mux.Handle("/register", handler.RegisterHandler(*registerService))
	mux.Handle("/user", getUserHandler)
	mux.Handle("/login", handler.LoginHandler(*loginService))
	mux.Handle("/upload", handler.UploadHandler(*musicService))
	mux.Handle("/static", getFileHandler)

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
