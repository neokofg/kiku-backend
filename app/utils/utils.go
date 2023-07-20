package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"kiku-backend/domain"
	"net/http"
	"strings"
	"time"
)

var jwtKey = []byte("your_secret_key")

func GenerateJWT(user domain.User) (string, error) {
	expirationTime := time.Now().Add(8640 * time.Hour) // Token valid for 24 hours; adjust as needed

	claims := &jwt.StandardClaims{
		Subject:   user.Login,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(r *http.Request) (string, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return "", fmt.Errorf("no token provided")
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	claims := &jwt.StandardClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", fmt.Errorf("invalid token signature")
		}
		return "", fmt.Errorf("could not parse token")
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	return claims.Subject, nil // Return the user's login
}
