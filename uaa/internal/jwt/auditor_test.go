package jwt

import (
	"testing"
)

func TestToken(t *testing.T) {
	t.Run("it should generate and validate token successfully", func(t *testing.T) {
		auditor := NewAuditor(
			map[string]string{
				"public":  "sample_public.pem",
				"private": "sample_private.pem",
			},
			EXPIRE_SECOND,
		)
		tokenString := auditor.GenerateToken(Claims{
			Scopes: []string{"read", "delete"},
		})

		claims, ok := auditor.ParseToken(tokenString)
		if !ok {
			t.Errorf("expected token string to be valid but not")
		}

		if claims.Scopes[0] != "read" && claims.Scopes[1] != "delete" {
			t.Errorf("expected decoded scopes to be equal to source")
		}
	})

	t.Run("it should fail when validating expired token", func(t *testing.T) {
		auditor := NewAuditor(
			map[string]string{
				"public":  "sample_public.pem",
				"private": "sample_private.pem",
			},
			-EXPIRE_SECOND,
		)
		tokenString := auditor.GenerateToken(Claims{})
		if _, ok := auditor.ParseToken(tokenString); ok {
			t.Errorf("expected token string to be invalid but valid")
		}
	})
}
