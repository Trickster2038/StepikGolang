package main

import (
	"fmt"
)

func main() {
	var a float64 = 0
	fmt.Scan(&a)
	if a > 10000 {
		fmt.Printf("%e", a)
	} else if a < 0 {
		fmt.Printf("число %.2f не подходит", a)
	} else {
		// a = a * a
		fmt.Printf("%.4f", a*a)
	}
}
