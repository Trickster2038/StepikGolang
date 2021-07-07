package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, y_cp, i, x_cp int = 0, 0, 0, 0, 0
	fmt.Scan(&x)
	fmt.Scan(&y)
	i = 0
	x_cp = x

	for x_cp > 0 {
		x_cp /= 10
		i++
	}
	i--

	for i >= 0 {
		y_cp = y
		for y_cp > 0 {
			if (y_cp % 10) == ((x / int(math.Round(math.Pow(10, float64(i))))) % 10) {
				fmt.Println((x / int(math.Round(math.Pow(10, float64(i))))) % 10)
			}
			y_cp /= 10
		}
		i--
	}
}
