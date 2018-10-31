package main

import "fmt"

// Create your own type. Have the underlying type be an int.
type hotdog int

var x hotdog

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v\n", x)
}
