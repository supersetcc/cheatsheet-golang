package main

import "fmt"

const hello = 0         // untyped
const hello_int int = 0 // typed
const (
	a int     = 42
	b float64 = 42.78
	c string  = "james"
)

func main() {
	fmt.Println(hello)
	fmt.Println(hello_int)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
