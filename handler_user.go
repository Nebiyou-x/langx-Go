package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Nebiyou-x/Golang/internal/database"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (apiCfg *apiConfig) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"` // Fixed struct tag
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	userID := uuid.New()
	now := time.Now().UTC()

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        pgtype.UUID{Bytes: userID, Valid: true},
		CreatedAt: pgtype.Timestamp{Time: now, Valid: true},
		UpdatedAt: pgtype.Timestamp{Time: now, Valid: true},
		Name:      params.Name,
	})

	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondwithJSON(w, 200, user) // 201 Created is more appropriate
}
