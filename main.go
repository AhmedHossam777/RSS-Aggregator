package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

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
	v1Router.Get("/users", apiCfg.handlerGetUser)
	v1Router.Post("/feeds", apiCfg.handlerCreateFeed)

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
