package main

import (
	"fmt"
)

func main() {
	for i:=10; i<=100; i++ {
		if i%4 == 0 {
			fmt.Printf("0 reminder, %v\n", i%4)
		} else if i%4 == 2 {
			fmt.Printf("2 reminder, %v\n", i%4)
		} else {
			fmt.Printf("useless %v\n", i%4)
		}
	}
}
