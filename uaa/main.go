package main

import (
	"flag"
	_ "net/http/pprof"
	"uaa/internal/jwt"
)

type AppContext struct {
	Data    Data
	Auditor jwt.Auditor
}

func main() {
	port := flag.String("port", ":8080", "server port")
	accounts := flag.String("accounts", "accounts.json", "account data")
	privateKey := flag.String("privateKey", "keys/private.pem", "private key path")
	publicKey := flag.String("publicKey", "keys/public.pem", "public key path")
	tokenExpiresSecond := flag.Duration(
		"tokenExpiresSecond",
		jwt.EXPIRE_SECOND,
		"token expires in second",
	)
	flag.Parse()

	ctx := AppContext{
		Data: readAccounts(*accounts),
		Auditor: jwt.NewAuditor(
			map[string]string{
				"private": *privateKey,
				"public":  *publicKey,
			},
			*tokenExpiresSecond,
		),
	}

	ctx.newServer(map[string]string{
		"port": *port,
	})
}
