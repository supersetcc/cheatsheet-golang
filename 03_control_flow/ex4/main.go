package main

import (
	"fmt"
)

func main() {
	year := 1989
	for {
		fmt.Println(year)
		year++
		if year > 2018 {
			break
		}
	}
}
