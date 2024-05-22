package main

import (
	"net/http"
	"strings"

	"github.com/thegouge/blog-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			respondWithError(w, 400, "you need to supply an auth token!")
			return
		}

		s := strings.Split(authHeader, " ")

		if s[0] != "ApiKey" || len(s) != 2 {
			respondWithError(w, 400, "invalid auth token!")
			return
		}

		user, err := cfg.DB.GetUserByAPI(r.Context(), s[1])

		if err != nil {
			respondWithError(w, 404, "Error finding user with the provided API Key")
			return
		}

		handler(w, r, user)
	}
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
