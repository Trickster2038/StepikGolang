package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := strings.ReplaceAll(scanner.Text(), ",", ".")
	s2 := strings.ReplaceAll(s, " ", "")

	words := strings.Split(s2, ";")

	a, _ := strconv.ParseFloat(words[0], 64)
	b, _ := strconv.ParseFloat(words[1], 64)

	fmt.Printf("%.4f", a/b)
}
