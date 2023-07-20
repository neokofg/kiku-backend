package routes

import (
	"database/sql"
	"github.com/rs/cors"
	"kiku-backend/app/application"
	"kiku-backend/app/handler"
	"kiku-backend/infrastructure"
	"log"
	"net/http"
)

func InitializeRoutes(db *sql.DB) {
	userRepo := &infrastructure.SqlUserRepository{DB: db}
	musicRepo := &infrastructure.SqlMusicRepository{DB: db}
	registerService := &application.RegisterService{Repo: userRepo}
	loginService := &application.LoginService{Repo: userRepo}
	musicService := &application.MusicService{Repo: musicRepo, UserRepository: userRepo}
	getUserHandler := &handler.GetUserHandler{LoginService: loginService}

	mux := http.NewServeMux()
	mux.Handle("/register", handler.RegisterHandler(*registerService))
	mux.Handle("/user", getUserHandler)
	mux.Handle("/login", handler.LoginHandler(*loginService))
	mux.Handle("/upload", handler.UploadHandler(*musicService))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // your app's origin
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})

	corsHandler := c.Handler(mux)

	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
