package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	for i := 0; i < len(a); i++ {
		if strings.Count(a, string(a[i])) > 1 {
			a = strings.Replace(a, string(a[i]), "", -1)
		}
		// fmt.Println(i)
	}
	fmt.Println(a)
	// put your code here
}
