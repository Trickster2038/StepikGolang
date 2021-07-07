package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	rs := []rune(text)
	// -2 in online IDE
	if unicode.IsUpper(rs[0]) && string(rs[len(rs)-2]) == "." {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
	fmt.Println(string(rs[len(rs)-2]))
}
