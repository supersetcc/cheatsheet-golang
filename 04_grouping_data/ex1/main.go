package main

import (
	"fmt"
)

func main() {
	x := [5]int{}
	for i := 0; i < 5; i++ {
		x[i] = i
	}
	fmt.Println(x)

	y := [5]int{1, 2, 3, 4, 5}
	for i, v := range y {
		fmt.Println(i, v)
	}
}
