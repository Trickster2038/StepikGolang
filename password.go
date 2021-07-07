package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var a string
	// fmt.Scan(&a)

	a, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	a = strings.TrimSuffix(a, "\n")

	rs := []rune(a)

	fl := true
	if len(rs) < 5 {
		fl = false
	}

	for i := range rs {
		if !(unicode.IsDigit(rs[i]) || unicode.Is(unicode.Latin, rs[i])) {
			fl = false
		}
	}

	if fl {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
	// unicode.IsDigit(r) || unicode.IsLetter(r)
	// put your code here
}
