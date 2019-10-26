package main

import "fmt"

func main() {
	// hello world
	fmt.Println("hello world")

	// multiple return values
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

func swap(a, b string) (string, string) {
	return b, a
}
