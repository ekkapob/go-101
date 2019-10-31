package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	/*** Display message ***/
	fmt.Println("Hello World!")
	fmt.Println("---------------------1")

	/*** Function in Go ***/
	fmt.Println(areaOfCircle(5.234))
	fmt.Println("---------------------2")

	/*** Function can return multilple values ***/
	fmt.Println("swapping text", "result is")
	fmt.Println(swapText("hello", "world"))
	fmt.Println("---------------------3")

	/*** Variables ***/
	var x string // default value for string is ""
	x = "hello"
	// this is equivalent to 2 lines above for creating x and assigning a value to x
	y := "world"
	fmt.Println(x, y)
	fmt.Println("---------------------4")

	/*** String concat ***/
	fmt.Println("Hello" + "World")
	fmt.Println("---------------------5")

	/*** Struct ***/
	// see below [struct]
	p := Person{}                        // create a variable p of type Person struct
	fmt.Println(p.Firstname, p.Lastname) // default values for Firstname and Lastname are ""

	p = Person{
		Firstname: "John",
		Lastname:  "Dale",
	} // create a new struct and assigning values
	fmt.Println(p.Firstname, p.Lastname)
	fmt.Println("---------------------6")

	/*** Struct Methods ***/
	// See [struct-meth]
	fmt.Println(p.fullName())
	fmt.Println("---------------------7")

	/*** Method ***/
	// ref: https://tour.golang.org/methods/3
	// see below [meth]
	var m MyInt
	m = 100 // shorter version is m := MyInt(100)
	fmt.Println(m.double())
	fmt.Println("---------------------8")

	/*** Pointer ***/
	q := "hello"
	var pt *string // how you create pointer variable (default is nil)
	// pt := new(string) is another way to create a pointer variable
	pt = &q // assign pointer (pt) to address of q (&q)
	fmt.Println("value of pt is an address (q's address):", pt)
	fmt.Println("address of q:", &q)
	fmt.Println("value of data pt is pointing to (q):", *pt)
	*pt = "world" // modify data pt is point to (modifying q data)
	fmt.Println("value of q is modified to:", q)
	fmt.Println("---------------------9")

	/*** Struct Pointer ***/
	pe := &Person{
		Firstname: "Jane",
		Lastname:  "Doe",
	}
	fmt.Println("<<<with struct receiver>>>")
	fmt.Println("default kids are", pe.Kids)
	// call a method with pointer receiver, so source struct will be modified
	pe.addKids(2)
	// pe is not impacted because it is struct receiver not pointer reciever
	fmt.Println("modified kids are", pe.Kids)
	fmt.Println("<<<with stuct pointer receiver>>>")
	fmt.Println("default age is", pe.Age)
	// call a method with pointer receiver, so source struct will be modified
	pe.addAge(30)
	fmt.Println("modified age is", pe.Age)
	fmt.Println("---------------------10")

	/*** Array ***/
	items := [2]string{"hello", "world"}
	modifyArray(items, "!")
	fmt.Println("list is", items)
	fmt.Println("---------------------11")

	/*** Slice ***/
	s := []string{"hello", "world"}
	modifySlice(s, "!")
	fmt.Println("list is", s)
	fmt.Println("---------------------12")

	/*** Map ***/
	var mm map[string]string
	mm = map[string]string{"a": "hello"}
	fmt.Println("map was", mm["a"])
	modifyMap(mm, "hi")
	fmt.Println("map is modified to", mm["a"])
	fmt.Println("---------------------13")

	/*** Go routine ***/
	fmt.Println("start go routines 1")
	go doSomething()
	time.Sleep(time.Second * 2) // sleep to make sure a routine is done
	fmt.Println("finish go routines")
	fmt.Println("---------------------14")

	fmt.Println("start go routines 2")
	done := make(chan bool)
	go doSomethingWithChan(done)
	<-done
	fmt.Println("finish go routines")
	fmt.Println("---------------------15")

	fmt.Println("start go routines 3")
	go doSomethingWithChan(done)
	select {
	case <-done:
		fmt.Println("finish go routines")
	}
	fmt.Println("---------------------16")

	/*** Go routine with select ***/
	c := make(chan int)
	i := 99
	go doSomethingSelect(c, done)
	for {
		select {
		case c <- i:
			i += 1
		case <-done:
			// goto to break for loop in main
			goto exit
		}
	}

exit:
	fmt.Println("---------------------17")
}

func modifyArray(list [2]string, s string) {
	list[1] = s
}

func modifySlice(slice []string, s string) {
	slice[1] = s
}

func modifyMap(m map[string]string, s string) {
	m["a"] = s
}

// [struct]
// This is how you create a struct
type Person struct {
	Firstname string
	Lastname  string
	Age       int
	Kids      int
}

// [struct-meth]
// Method is function with a receiver (Golang calls "Receiver Function")
// a receiver is (p Person) below
// so, this is similar to class method in other OOP
// ref: https://appdividend.com/2019/03/23/golang-receiver-function-tutorial-go-function-receivers-example/
func (p Person) fullName() string {
	return p.Firstname + " " + p.Lastname
}

func (p Person) addKids(kids int) {
	p.Kids = kids
}

func (p *Person) addAge(age int) {
	p.Age = age
}

// [meth]
type MyInt int

// a method (with MyInt receiver)
func (m MyInt) double() MyInt {
	return m * m
}

func areaOfCircle(r float32) float32 {
	return math.Pi * r * r
}

func swapText(a string, b string) (string, string) {
	return b, a
}

func doSomething() {
	fmt.Println("do something")
}

func doSomethingWithChan(done chan bool) {
	fmt.Println("do something")
	done <- true
}

func doSomethingSelect(c chan int, done chan bool) {
	fmt.Println("1. wait for c", <-c)
	fmt.Println("2. wait for c", <-c)
	fmt.Println("3. wait for c", <-c)
	fmt.Println("4. wait for c", <-c)
	done <- true
}
