package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github/AhmedHossam777/RSS-Aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(
	w http.ResponseWriter,
	r *http.Request,
	user database.User,
) {
	feedId := chi.URLParam(r, "id")
	parsedFeedId, err := uuid.Parse(feedId)
	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("invalid uuid format: %v", err))
		return
	}

	feedFollow, err := apiCfg.db.CreateFeedFollow(
		r.Context(), database.CreateFeedFollowParams{
			ID:     uuid.New(),
			UserID: user.ID,
			FeedID: parsedFeedId,
		},
	)

	if err != nil {
		ResponseWithError(
			w, 400, fmt.Sprintf("Error creating feed follow: %v", err),
		)
		return
	}

	ResponseWithJson(w, 201, DatabaseFeedFollowToFeedFollow(feedFollow))
}
