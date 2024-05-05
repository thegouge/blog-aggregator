package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/thegouge/blog-aggregator/internal/database"
)

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

	feedObject, err := api.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error creating feed object: %v", err))
		return
	}

	feedFollowObject, err := api.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
		FeedID:    feedObject.ID,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error creating feed object: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, struct {
		Feed       Feed
		FeedFollow FeedFollow
	}{
		Feed:       databaseFeedToFeed(feedObject),
		FeedFollow: databaseFeedFollowToFeedFollow(feedFollowObject),
	})
}

func (api *apiConfig) HandleGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := api.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, 500, "Something went wrong fetching all the feeds!")
	}

	respondWithJSON(w, http.StatusOK, feeds)
}
