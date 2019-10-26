package jwt

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const EXPIRE_SECOND = 60

type Auditor struct {
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
	PublicKey []byte
	ExpiresIn time.Duration
}

type Claims struct {
	Scopes []string `json:"scopes"`
	jwt.StandardClaims
}

func NewAuditor(filePaths map[string]string, expiresIn time.Duration) Auditor {
	privateKey, err := ioutil.ReadFile(filePaths["private"])
	if err != nil {
		log.Fatalf("read private key: %v", err)
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		log.Fatalf("parse RSA private key: %v", err)
	}
	publicKey, err := ioutil.ReadFile(filePaths["public"])
	if err != nil {
		log.Fatalf("read public key: %v", err)
	}
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		log.Fatalf("parse RSA private key: %v", err)
	}
	return Auditor{
		signKey:   signKey,
		verifyKey: verifyKey,
		PublicKey: publicKey,
		ExpiresIn: expiresIn * time.Second,
	}
}

func (a Auditor) GenerateToken(c Claims) string {
	if c.StandardClaims.ExpiresAt == 0 {
		c.StandardClaims.ExpiresAt = time.Now().Add(a.ExpiresIn).Unix()
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), c)
	tokenString, err := token.SignedString(a.signKey)
	if err != nil {
		log.Println("signing string: ", err)
		return ""
	}
	return tokenString
}

func (a Auditor) ParseToken(tokenString string) (*Claims, bool) {
	c := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, c,
		func(token *jwt.Token) (interface{}, error) {
			return a.verifyKey, nil
		})
	if err != nil {
		log.Println("verifying token string:", err)
		return nil, false
	}
	return c, token.Valid
}
