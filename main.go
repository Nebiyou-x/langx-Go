package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")
	portstring := os.Getenv("PORT")

	if portstring == "" {
		log.Fatal("Error occured")
	}

	router := chi.NewRouter()

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
