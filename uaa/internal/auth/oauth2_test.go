package auth

import (
	"fmt"
	"testing"
)

func TestOauthToken(t *testing.T) {
	t.Run("it should return bearer token", func(t *testing.T) {
		tokenString := "some-token"
		token, err := OauthToken(fmt.Sprint("Bearer ", tokenString))
		if err != nil {
			t.Errorf("expected valid token but not")
		}
		if token != tokenString {
			t.Errorf("expected token to be equal to input token")
		}
	})

	t.Run("invalid header", func(t *testing.T) {
		tokenString := "some-token"

		t.Run("no space between bearer and token", func(t *testing.T) {
			_, err := OauthToken(fmt.Sprint("Bearer", tokenString))
			if err == nil {
				t.Errorf("expected invalid token but valid")
			}
		})

		t.Run("no bearer", func(t *testing.T) {
			_, err := OauthToken(fmt.Sprint("Basic ", tokenString))
			if err == nil {
				t.Errorf("expected invalid token but valid")
			}
		})
	})
}
