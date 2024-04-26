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

func (api *apiConfig) addFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type feedParams struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := feedParams{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error decoding feed Object: %v", err))
		return
	}

	newId := uuid.NewString()

	feedData := database.CreateFeedParams{
		ID:        newId,
		CreatedAt: time.Now(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	}

	feedObject, err := api.DB.CreateFeed(r.Context(), feedData)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error creating feed object: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, feedObject)
}

func (api *apiConfig) HandleGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := api.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, 500, "Something went wrong fetching all the feeds!")
	}

	respondWithJSON(w, http.StatusOK, feeds)
}

func (api *apiConfig) addFeedFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type requestParams struct {
		FeedId string `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := requestParams{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error decoding json body: %v", err))
		return
	}

	newId := uuid.NewString()

	feedFollowData := database.CreateFeedFollowParams{
		ID:        newId,
		CreatedAt: time.Now(),
		FeedID:    params.FeedId,
		UserID:    user.ID,
	}

	feedFollowObject, err := api.DB.CreateFeedFollow(r.Context(), feedFollowData)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error creating feed object: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, feedFollowObject)
}
