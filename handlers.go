package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/thegouge/blog-aggregator/internal/database"
)

func healthHandler(w http.ResponseWriter, Request *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func errHandler(w http.ResponseWriter, Request *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}

func (api *apiConfig) HandleUserCreate(w http.ResponseWriter, r *http.Request) {
	type createUserRequestBody struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := createUserRequestBody{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error decoding User Name: %v", err))
		return
	}

	newId := uuid.NewString()

	nowTime := time.Now()

	newUser, err := api.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        newId,
		CreatedAt: nowTime,
		UpdatedAt: nowTime,
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error creating new user %s, %v", params.Name, err))
		return
	}

	respondWithJSON(w, http.StatusOK, newUser)
}

func (api *apiConfig) HandleGetUserByAPI(w http.ResponseWriter, r *http.Request) {
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

	user, err := api.DB.GetUserByAPI(r.Context(), s[1])

	if err != nil {
		respondWithError(w, 404, "Error finding user with the provided API Key")
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}
