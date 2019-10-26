package auth

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestDecodeBasic(t *testing.T) {
	t.Run("it should return valid username and password", func(t *testing.T) {
		username := "john"
		password := "doe"
		auth := basicAuth(username, password)
		decodedUsername, decodedPassword, err := DecodeBasic(auth)
		if decodedUsername != username || decodedPassword != password || err != nil {
			t.Error(err)
		}
	})

	t.Run("invalid auth header", func(t *testing.T) {
		t.Run("invalid basic auth header", func(t *testing.T) {
			token := base64.StdEncoding.EncodeToString([]byte(fmt.Sprint("john:doe")))
			_, _, err := DecodeBasic(fmt.Sprintf("Invalid %v", token))
			if err == nil {
				t.Error(err)
			}
		})

		t.Run("invalid username password pattern", func(t *testing.T) {
			token := base64.StdEncoding.EncodeToString([]byte(fmt.Sprint("johndoe")))
			_, _, err := DecodeBasic(fmt.Sprintf("Basic %v", token))
			if err == nil {
				t.Error(err)
			}
		})
	})

}

func basicAuth(username, password string) string {
	token := base64.StdEncoding.EncodeToString([]byte(fmt.Sprint(username, ":", password)))
	return fmt.Sprintf("Basic %v", token)
}
