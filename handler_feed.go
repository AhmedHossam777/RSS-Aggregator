package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github/AhmedHossam777/RSS-Aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeed(
	w http.ResponseWriter, r *http.Request, user database.User,
) {

	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	param := parameters{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&param)
	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("error parsing json: %v", err))
		return
	}

	feed, err := apiCfg.db.CreateFeed(
		r.Context(), database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      param.Name,
			Url:       param.URL,
			UserID:    user.ID,
		},
	)

	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("error creating feed: %v", err))
		return
	}

	ResponseWithJson(w, 201, DatabaseFeedToFeed(feed))
}
