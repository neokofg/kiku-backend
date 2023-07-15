package handler

import (
	"encoding/json"
	"kiku-backend/app/application"
	"kiku-backend/app/utils"
	"net/http"
)

type GetUserHandler struct {
	LoginService *application.LoginService
}

func (h *GetUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Validate token and extract the user's login
	login, err := utils.ValidateToken(r)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Fetch user data
	user, err := h.LoginService.GetUserByLogin(login)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Respond with the user data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
