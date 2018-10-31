package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	// to figure out what % sign means.. https://godoc.org/fmt
	s := fmt.Sprintf("%#v\t%#v\t%#v", x, y, z)
	fmt.Println(s)
}
