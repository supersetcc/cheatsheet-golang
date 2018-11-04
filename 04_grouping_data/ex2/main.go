package main

import (
	"fmt"
)

func main() {
	x := make([]int, 10, 10)
	for i:=0; i<10; i++ {
		x[i] = i
	}

	for i,v := range x {
		fmt.Println(i, v)
	}
}
