package main

import (
	"fmt"
	"net/http"

	"github/AhmedHossam777/RSS-Aggregator/internal/auth"
	"github/AhmedHossam777/RSS-Aggregator/internal/database"
)

type authHandler func(
	w http.ResponseWriter, r *http.Request, user database.User,
)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the api key
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			ResponseWithError(
				w, 403, fmt.Sprintf("error getting the api key, %v", err),
			)
			return
		}

		// now get the user of that api key
		user, err := apiCfg.db.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			ResponseWithError(w, 400, fmt.Sprintf("error fetching the user, %v", err))
			return
		}

		handler(w, r, user)
	}
}
