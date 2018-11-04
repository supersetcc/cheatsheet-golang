package main

import (
	"fmt"
)

func main() {
	x := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	x[`hunjae`] = []string{`alice`} // ex 9
	delete(x, `no_dr`)  // ex 10
	for k, v := range x {
		fmt.Println(k)
		for i, favorite := range v {
			fmt.Println("\t", i, favorite)
		}
	}
}
