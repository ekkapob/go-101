HOST = http://localhost:8080

all: dev

dev:
	realize start

test:
	go test -cover -tags test

coverprofile:
	go test -coverprofile cp.out -tags test
	go tool cover -html=cp.out
	rm cp.out

pprof:
	go build
	example-3 &
	echo "GET http://localhost:8080/authen/pass" | vegeta attack -rate=2000 -duration=10s | vegeta report &
	go tool pprof example-3 http://localhost:8080/debug/pprof/profile
	# go tool pprof -http=:6060 <part to pprof.samples.cpu.005.pb.gz>"

e2e:
	HOST=$(HOST) mocha

.PHONY: all test clean
