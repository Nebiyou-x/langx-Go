package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")
	portstring := os.Getenv("PORT")

	if portstring == "" {
		log.Fatal("Error occured")
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

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portstring,
	}

	log.Printf("server starting on port %v", portstring)

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
