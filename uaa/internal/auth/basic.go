package auth

import (
	"encoding/base64"
	"errors"
	"strings"
)

func DecodeBasic(authHeader string) (username, password string, err error) {
	tokens := strings.Split(authHeader, " ")
	if strings.ToLower(tokens[0]) != "basic" {
		return username, password, errors.New("expected basic auth header")
	}

	encodedCredentials, err := base64.StdEncoding.DecodeString(tokens[1])
	credentials := strings.Split(string(encodedCredentials), ":")
	if len(credentials) != 2 {
		return username, password, errors.New("expected username and password")
	}

	return credentials[0], credentials[1], nil
}
