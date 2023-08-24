package http

import (
	"encoding/json"
	"log"
	"net/http"
)

func (d *Delivery) access(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userId := r.URL.Query().Get("user_id")
	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("invalid user_id param")
		return
	}

	tokens, err := d.user.GenerateTokenPair(ctx, userId)
	if err != nil {
		log.Printf("user service: generate token pair: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("failed to generate token pair")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokens)
}

func (d *Delivery) refresh(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userId := r.URL.Query().Get("user_id")
	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("invalid user_id param")
		return
	}

	refreshToken := r.URL.Query().Get("refresh_token")
	if refreshToken == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("invalid refresh_token param")
		return
	}

	tokens, err := d.user.RefreshTokenPair(ctx, userId, refreshToken)
	if err != nil {
		log.Printf("user service: refresh token pair: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("failed to refresh tokens")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokens)
}
