package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Nebiyou-x/Golang/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load(".env")
	portstring := os.Getenv("PORT")

	if portstring == "" {
		log.Fatal("Error occured , PORT not found")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("Error occured , DB_URL not found")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("cant connect to database", err)
	}

	dbQueries := database.New(conn)

	apiCfg := apiConfig{
		DB: database.New(db),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	//v1Router.HandleFunc("/healthz", HandleReadiness)
	v1Router.Get("/healthz", HandleReadiness)
	v1Router.Get("/err", HandlErr)
	v1Router.Post("/users", apiCfg.HandleCreateUser)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portstring,
	}

	log.Printf("server starting on port %v", portstring)

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
