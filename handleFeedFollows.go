package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/thegouge/blog-aggregator/internal/database"
)

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

	feedFollowObject, err := api.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        newId,
		CreatedAt: time.Now(),
		FeedID:    params.FeedId,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error creating feed object: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, feedFollowObject)
}

func (api *apiConfig) deleteFeedFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowID := r.PathValue("feedFollowID")

	err := api.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{UserID: user.ID, ID: feedFollowID})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete feed follow")
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}

func (api *apiConfig) getUsersFeedFollowsHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	userFeeds, err := api.DB.GetUsersFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't find feeds for the current user")
		return
	}

	respondWithJSON(w, http.StatusOK, userFeeds)
}
