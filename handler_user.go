package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github/AhmedHossam777/RSS-Aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(
	w http.ResponseWriter, r *http.Request,
) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("error parsing json: %v", err))
		return
	}

	user, err := apiCfg.db.CreateUser(
		r.Context(), database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      params.Name,
		},
	)

	createdUser := DatabaseUserToUser(user)

	if err != nil {
		ResponseWithError(w, 500, fmt.Sprintf("error creating new user: %v", err))
		return
	}

	ResponseWithJson(w, 200, createdUser)
}
