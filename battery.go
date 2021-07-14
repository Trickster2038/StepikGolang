package main

import (
	"fmt"
	"strconv"
)

type Battery int

func (battery Battery) String() string {
	x := 0
	del := 1
	//fmt.Println(strconv.ParseInt("02", 10, 64))
	for i := 1; i < 11; i++ {
		x += (int(battery) / del) % 10
		del *= 10
	}
	rs := []rune("[----------]")
	rs[0] = '['
	for i := 1; i < 11-x; i++ {
		rs[i] = ' '
	}
	for i := 11 - x; i < 11; i++ {
		rs[i] = 'X'
	}
	rs[11] = ']'
	return fmt.Sprintf("(%d) %s", x, string(rs))
}

func main() {
	var a string
	// fmt.Scan(&a)
	a = "0011"
	n, _ := strconv.ParseInt(a, 10, 64)
	batteryForTest := Battery(n)
	fmt.Println(batteryForTest)
}
