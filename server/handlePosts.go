package main

import (
	"net/http"
	"strconv"

	"github.com/thegouge/blog-aggregator/internal/database"
)

func (api *apiConfig) HandleGetPostsByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	var resultLimit int32 = 10
	queriedLimit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)

	if err == nil {
		resultLimit = int32(queriedLimit)
	}

	userFollowedPosts, err := api.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		ID:    user.ID,
		Limit: resultLimit,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't find posts for the current user")
		return
	}

	respondWithJSON(w, http.StatusOK, userFollowedPosts)
}
