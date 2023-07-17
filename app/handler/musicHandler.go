package handler

import (
	"io"
	"kiku-backend/app/application"
	"kiku-backend/app/utils"
	"kiku-backend/domain"
	"net/http"
	"os"
)

type uploadRequest struct {
	Name   string      `json:"name"`
	Author string      `json:"author"`
	User   domain.User `json:"user"`
	File   string      `json:"file"`
}

func UploadHandler(s application.MusicService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req uploadRequest

		// Parse the incoming request containing the form data
		err := r.ParseMultipartForm(10 << 20) // limit your maxMultipartMemory
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Get a reference to the fileHeaders.
		file, header, err := r.FormFile("uploadfile")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Get and print out file information.
		req.Name = header.Filename
		req.Author = r.FormValue("author")

		// Get user login from token
		login, err := utils.ValidateToken(r)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Fetch the user from the database
		// Assuming the existence of a GetUserByLogin function in your music service
		user, err := s.GetUserByLogin(login)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		req.User = *user
		// Create the file in the local file system
		dst, err := os.Create("/projects/kiku-backend/static/" + req.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Build the URL to the file
		req.File = "http://localhost:8080/static/" + req.Name

		// Convert your uploadRequest to domain.Music
		music := domain.Music{
			Name:   req.Name,
			Author: req.Author,
			User:   req.User.ID, // assuming your User domain model has an ID field
			Url:    req.File,
		}
		// Use the Create method to save the song
		err = s.UploadMusic(&music)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Return a response
		w.Write([]byte("Successfully Uploaded File\n"))
	}
}
func GetFileHandler(w http.ResponseWriter, r *http.Request) {
	// Получите URL файла из параметра запроса
	fileURL := r.URL.Query().Get("file")

	// Загрузите файл по его URL
	response, err := http.Get(fileURL)
	if err != nil {
		http.Error(w, "Failed to fetch file", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	// Установите заголовки ответа для указания типа контента и длины файла
	w.Header().Set("Content-Type", response.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", response.Header.Get("Content-Length"))

	// Скопируйте содержимое файла в тело ответа
	_, err = io.Copy(w, response.Body)
	if err != nil {
		http.Error(w, "Failed to copy file content", http.StatusInternalServerError)
		return
	}
}
