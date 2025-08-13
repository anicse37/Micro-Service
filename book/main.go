package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret_key") // must match auth service

func validateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	books := []string{"Book A", "Book B"}
	json.NewEncoder(w).Encode(books)
}

func main() {
	http.HandleFunc("/books", validateJWT(getBooks))
	fmt.Println("Books service running on 8082")
	http.ListenAndServe(":8082", nil)
}
