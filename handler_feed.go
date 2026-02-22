package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github/AhmedHossam777/RSS-Aggregator/internal/auth"
	"github/AhmedHossam777/RSS-Aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeed(
	w http.ResponseWriter, r *http.Request,
) {
	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		ResponseWithError(w, 403, fmt.Sprintf("error getting the api key, %v", err))
		return
	}

	user, err := apiCfg.db.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("error fetching the user, %v", err))
		return
	}

	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	param := parameters{}
	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&param)
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
			UserID: uuid.NullUUID{
				UUID:  user.ID,
				Valid: true,
			},
		},
	)

	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("error creating feed: %v", err))
		return
	}

	ResponseWithJson(w, 201, feed)
}
