package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func adding(a, b string) int64 {
	//var a, b string
	//fmt.Scan(&a)
	//fmt.Scan(&b)
	runes1 := []rune(a)
	runes2 := []rune(b)
	//fmt.Println(runes1)
	str1 := ""
	str2 := ""
	for _, elem := range runes1 {
		//fmt.Println(elem)
		if unicode.IsDigit(rune(elem)) {
			str1 = str1 + string(elem)
		}
	}
	for _, elem := range runes2 {
		if unicode.IsDigit(rune(elem)) {
			str2 = str2 + string(elem)
		}
	}
	//fmt.Println(str1)
	//a1, _ := strconv.ParseInt(str1, 0, 0)
	//b1, _ := strconv.ParseInt(str2, 0, 0)
	a1, _ := strconv.Atoi(str1)
	b1, _ := strconv.Atoi(str2)
	return int64(a1 + b1)
}

func main() {
	x := "a13hhh"
	y := "aa45gg"
	z := adding(x, y)
	fmt.Println(strconv.Atoi("23"))
	fmt.Println(z)
}
