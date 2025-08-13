package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret_key") // use env var in prod

func login(w http.ResponseWriter, r *http.Request) {
	var creds struct{ Username, Password string }
	json.NewDecoder(r.Body).Decode(&creds)

	// Fake check
	if creds.Username != "admin" || creds.Password != "pass" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": creds.Username,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
	})
	tokenString, _ := token.SignedString(jwtKey)
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func main() {
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8081", nil)
}
