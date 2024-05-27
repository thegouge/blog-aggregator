package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/thegouge/blog-aggregator/internal/database"
)

type UserRequestBody struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (api *apiConfig) HandleUserCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := UserRequestBody{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error decoding User Params: %v", err))
		return
	}

	newId := uuid.NewString()

	nowTime := time.Now()

	newUser, err := api.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        newId,
		CreatedAt: nowTime,
		Name:      params.Name,
		Crypt:     params.Password,
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

func (api *apiConfig) HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := UserRequestBody{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error decoding User Params: %v", err))
		return
	}

	user, err := api.DB.LogInAsUser(r.Context(), database.LogInAsUserParams{
		Name:  params.Name,
		Crypt: params.Password,
	})

	if user.Name == "" {
		respondWithError(w, http.StatusForbidden, "Error logging in, username or password incorrect")
		return
	}

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error logging in to user %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}
