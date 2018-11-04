package main

import (
	"fmt"
)

func main() {
	// ex 4
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 42)
	fmt.Println(x)

	x = append(x, 53,54,55)
	fmt.Println(x)

	y := []int{56,57,58}
	x = append(x, y...)
	fmt.Println(x)

	// ex 5
	x = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y = append(x[:3], x[6:]...)
	fmt.Println(y)
}
