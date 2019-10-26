### Install Libraries

#### Realize

```
https://github.com/oxequa/realize
```

<<<<<<< HEAD
=======
with issue try about cli.v2


```
go get gopkg.in/urfave/cli.v2@master
```

### HTTP Benchmark Tools

```
$ go get github.com/adjust/go-wrk
$ go get github.com/tsenart/vegeta
$ brew install wrk
```

#### Graphviz

```
$ brew install graphviz
```

#### Profiling Steps

1. Build and run server

```
$ go build
$ ./example-3
```

2. Standby Profiling

```
$ go tool pprof --seconds=5 localhost:8080/debug/pprof/profile
```

3. Load testing
```
$ echo "GET http://localhost:8080/" | vegeta attack -rate=5000 -duration=1s | vegeta report

OR

$ go-wrk -d 60 http://localhost:8080

OR

$ wrk -t12 -c400 -d30s http://localhost:8080
```

4. Back to #2

Copy path to profile ex. `Saved profile in /Users/ek/pprof/pprof.samples.cpu.007.pb.gz`


```
(pprof) exit
```

5. Open PPROF on saved profile

```
go tool pprof -http=:6060 /Users/ek/pprof/pprof.samples.cpu.007.pb.gz

```

#### Benchmarking Profile

1. Write go test with `for i := 0; i < b.N; i++`. Check `handlers_test.go`

2. Run go test benchmark

```
$ go test -bench . -cpuprofile prof.cpu
```

3. Open PPROF on saved profile

```
$ go tool pprof -http=:6060 prof.cpu
```

#### Information

Ignore files from Test Coverage

```
+build !test

package main
...
```

To Run

```
$ go test -cover -tags test
```

#### Practices

1. Develop Handler for 404 (page not found) and its Unit test

   ex: Browsing to "/not-found-url", the server is showing "Not Found" and HTTP status of 404 Not Found

