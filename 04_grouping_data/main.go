package main

import (
	"fmt"
)

func main() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Printf("%T\n", x)

	x[0] = 1
	x[9] = 10

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Printf("%T\n", x)

	x = append(x, 123)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Printf("%T\n", x)

	a := []string{"A", "B", "C"}

	b := []string{"A", "B", "C"}

	c := [][]string{a, b}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	m := map[string]int{
		"a":    39,
		"miss": 27,
		"b":    28,
	}
	m["c"] = 34
	fmt.Println(m)

	delete(m ,"c")
	fmt.Println(m)

	v, ok_v := m["a"]
	fmt.Println(v)
	fmt.Println(ok_v)

	if v, ok := m["a"]; ok {
		fmt.Println("there is v", v)
	}

	k, ok_k := m["c"]
	fmt.Println(k)
	fmt.Println(ok_k)
}