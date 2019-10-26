//+build !test

package main

import (
	"flag"
	_ "net/http/pprof"
)

func main() {
	port := flag.String("port", ":8080", "server port")
	flag.Parse()

	newServer(map[string]string{
		"port": *port,
	})
}
