package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"uaa/internal/jwt"
)

func TestPublicKeyHandler(t *testing.T) {
	ctx := AppContext{
		Auditor: jwt.NewAuditor(
			map[string]string{
				"private": "keys/private.pem",
				"public":  "keys/public.pem",
			},
			60,
		),
	}
	url := "/oauth2/public_key"

	t.Run("it should return the public key with GET request", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			t.Error(err)
		}
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(ctx.PublicKeyHandler)
		handler.ServeHTTP(resp, req)

		if status := resp.Code; status != http.StatusOK {
			t.Errorf("expected %v but got %v", http.StatusOK, status)
		}

		publicKey, err := ioutil.ReadFile("keys/public.pem")
		if err != nil {
			t.Error(err)
		}
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		if string(publicKey) != string(respBody) {
			t.Errorf("expected response to be equal to public key")
		}
	})

	t.Run("it should return unauthorized with other request types", func(t *testing.T) {
		reqTypes := []string{http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete}

		for _, reqType := range reqTypes {
			t.Run(fmt.Sprintf("%v request", reqType), func(t *testing.T) {
				req, err := http.NewRequest(reqType, url, nil)
				if err != nil {
					t.Error(err)
				}
				resp := httptest.NewRecorder()
				handler := http.HandlerFunc(ctx.PublicKeyHandler)
				handler.ServeHTTP(resp, req)
				if status := resp.Code; status != http.StatusUnauthorized {
					t.Errorf("expected %v but got %v", http.StatusUnauthorized, status)
				}
			})
		}
	})

}
