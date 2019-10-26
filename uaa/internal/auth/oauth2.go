package auth

import (
	"errors"
	"strings"
)

func OauthToken(authHeader string) (string, error) {
	tokens := strings.Split(authHeader, " ")
	if len(tokens) != 2 || strings.ToLower(tokens[0]) != "bearer" {
		return "", errors.New("invalid token")
	}
	return tokens[1], nil
}
