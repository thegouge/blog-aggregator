package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/thegouge/blog-aggregator/internal/database"
)

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
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error creating new user %s, %v", params.Name, err))
		return
	}

	respondWithJSON(w, http.StatusOK, newUser)
}

func (api *apiConfig) HandleGetUserByAPI(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, user)
}
