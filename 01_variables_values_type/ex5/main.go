package main

import "fmt"

type hotdog int

var x hotdog
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 1
	fmt.Println(x)

	//x = int(x) // not possible
	y = int(x)
	fmt.Println(y)
}
