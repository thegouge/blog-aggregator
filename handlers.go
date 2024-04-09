package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func (api *apiConfig) HandleUserCreate(w http.ResponseWriter, Request *http.Request) {
	type createUserRequestBody struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(Request.Body)
	params := createUserRequestBody{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error decoding User Name: %v", err))
		return
	}

	newId := uuid.New()

	nowTime := time.Now()

	newUser, err := api.DB.CreateUser(Request.Context(), database.CreateUserParams{
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
