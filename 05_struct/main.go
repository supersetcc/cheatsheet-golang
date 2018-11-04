package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "alice",
		last:  "han",
	}

	p2 := person{
		first: "hunjae",
		last:  "jung",
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first)
	fmt.Println(p2.first)
}
