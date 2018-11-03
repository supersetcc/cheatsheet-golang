package main

import (
	"fmt"
)

func main() {
	x := 1
	switch {
	case x == 1:
		fmt.Println("a")
	default:
		fmt.Println("b")
	}

	favSport := "tennis"
	switch favSport {
	case "tennis":
		fmt.Println("a")
	default:
		fmt.Println("b")
	}

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!false)
}
