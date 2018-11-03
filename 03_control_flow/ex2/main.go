package main

import (
	"fmt"
)

func main() {
	x := 0x41
	for ; x <= 0x41+25; x++ {
		for i := 0; i < 3; i++ {
			fmt.Printf("%#U\n", x)
		}
	}
}
