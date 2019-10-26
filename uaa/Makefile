HOST = http://localhost:8080

all: dev

dev:
	realize start

test:
	go test ./... -cover

coverprofile:
	go test -coverprofile cp.out
	go tool cover -html=cp.out
	rm cp.out

e2e:
	HOST=$(HOST) mocha

.PHONY: all test clean
