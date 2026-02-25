package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github/AhmedHossam777/RSS-Aggregator/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	db *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load the dotenv file")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL url is not found in the environment")
	}
	conn, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	queries := database.New(conn)

	apiCfg := apiConfig{
		db: queries,
	}

	go startScraping(queries, 10, time.Minute)

	router := chi.NewRouter()
	router.Use(
		cors.Handler(
			cors.Options{
				MaxAge:           300,
				AllowedOrigins:   []string{"https://*", "http://*"},
				AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
				AllowedHeaders:   []string{"Link"},
				AllowCredentials: false,
			},
		),
	)

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", HandlerReadiness)
	v1Router.Get("/err", HandlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetAllFeeds)
	v1Router.Post(
		"/follow-feed/{id}", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow),
	)
	v1Router.Delete(
		"/follow-feed/{id}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow),
	)
	v1Router.Get(
		"/follow-feed", apiCfg.middlewareAuth(apiCfg.handlerGetFeedsFollow),
	)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on Port: %v", portString)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
