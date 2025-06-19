package main

import (
	"net/http"
)

func HandleReadiness(w http.ResponseWriter, r *http.Request) {
	respondwithJSON(w, 200, struct{}{})
}
