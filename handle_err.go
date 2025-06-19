package main

import (
	"net/http"
)

func HandlErr(w http.ResponseWriter, r *http.Request) {
	respondwithJSON(w, 400, "Something went wrong")
}
