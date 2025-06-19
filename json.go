package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondwithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with Error: ", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	respondwithJSON(w, code, errResponse{
		Error: msg,
	})
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON rsponse : %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)
}
