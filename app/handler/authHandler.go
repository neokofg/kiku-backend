package handler

import (
	"encoding/json"
	"kiku-backend/app/application"
	"kiku-backend/app/utils"
	"kiku-backend/domain"
	"net/http"
)

func RegisterHandler(s application.RegisterService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user domain.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = s.Register(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Generate JWT token
		token, err := utils.GenerateJWT(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		// Return the JWT token in the response
		w.Write([]byte(token))
	}
}

type loginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func LoginHandler(s application.LoginService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req loginRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := s.Login(req.Login, req.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// Generate JWT token
		token, err := utils.GenerateJWT(*user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		// Return the JWT token in the response
		w.Write([]byte(token))
	}
}
