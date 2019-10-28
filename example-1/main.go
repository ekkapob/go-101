package main

import (
	"fmt"
	"time"
)

func main() {
	// hello world
	fmt.Println("hello world")

	fmt.Println("-----------------------------")

	// multiple return values
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println("-----------------------------")

	// basic go routine
	go greeting("1", nil)
	bye("1")
	time.Sleep(100)

	fmt.Println("-----------------------------")

	x := make(chan bool)
	go greeting("2", x)
	select {
	case <-x:
		bye("2")
	}

	fmt.Println("-----------------------------")

	y := make(chan bool)
	go greeting("3", y)
	<-y
	bye("3")

	fmt.Println("-----------------------------")

	z := make(chan bool)
	go greeting("4", z)
	for {
		select {
		case result := <-z:
			bye(fmt.Sprintln("4", result))
			return
		}
	}
}

func swap(a, b string) (string, string) {
	return b, a
}

func greeting(id string, success chan bool) {
	fmt.Println("Greeting", id)
	success <- true
}

func bye(id string) {
	fmt.Println("Bye", id)
}
