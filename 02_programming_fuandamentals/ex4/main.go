package main

import "fmt"

func main() {
	v := 180
	fmt.Printf("%b, %d, %#x\n", v, v, v)
	v = v << 1
	fmt.Printf("%b, %d, %#x\n", v, v, v)
	v = v << 1
	fmt.Printf("%b, %d, %#x\n", v, v, v)
}
