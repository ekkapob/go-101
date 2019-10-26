package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"uaa/internal/auth"
	"uaa/internal/jwt"
)

func (c AppContext) PublicKeyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write(c.Auditor.PublicKey)
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
}

func (c AppContext) ValidateHandler(w http.ResponseWriter, r *http.Request) {

	type Response struct {
		ValidToken bool `json:"valid_token"`
	}

	if r.Method == http.MethodPost {
		w.Header().Set("Content-Type", "application/json")

		token, err := auth.OauthToken(r.Header.Get("Authorization"))
		if err != nil {
			responseJSON(w, http.StatusOK, Response{false})
			return
		}

		_, valid := c.Auditor.ParseToken(token)
		responseJSON(w, http.StatusOK, Response{valid})
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
}

func (c AppContext) TokenHandler(w http.ResponseWriter, r *http.Request) {

	type Response struct {
		AccessToken string        `json:"access_token,omitempty"`
		Scopes      []string      `json:"scopes,omitempty"`
		ExpiresIn   time.Duration `json:"expires_in,omitempty"`
		Error       string        `json:"error,omitempty"`
	}

	if r.Method == http.MethodPost {
		w.Header().Set("Content-Type", "application/json")

		username, password, err := auth.DecodeBasic(r.Header.Get("Authorization"))
		if err != nil {
			log.Println("decoded auth header:", err)
			responseJSON(w, http.StatusBadRequest, Response{Error: "invalid basic auth"})
			return
		}

		account, err := c.Data.GetAccount(username, password)
		if err != nil {
			log.Println("fetching an account:", err)
			responseJSON(
				w,
				http.StatusBadRequest,
				Response{Error: "invalid username or password"},
			)
			return
		}

		accessToken := c.Auditor.GenerateToken(jwt.Claims{Scopes: account.Scopes})
		json.NewEncoder(w).Encode(Response{
			AccessToken: accessToken,
			Scopes:      account.Scopes,
			ExpiresIn:   c.Auditor.ExpiresIn / time.Second,
		})

		return
	}

	w.WriteHeader(http.StatusUnauthorized)
}

func responseJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(v)
}
