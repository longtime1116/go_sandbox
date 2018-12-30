package main

import (
	"fmt"
	"strconv"
	"strings"
)

func add(x int, y int) int {
	return x + y
}

func AddOmitType(x, y int) int {
	return x + y
}

func Swap(x, y string) (string, string) {
	return y, x
}

func SplitDate(s string) (y, m, d int) {
	ary := strings.Split(s, "-")
	y, _ = strconv.Atoi(ary[0])
	m, _ = strconv.Atoi(ary[1])
	d, _ = strconv.Atoi(ary[2])
	return
}

func main() {
	fmt.Println("Function.", add(2, 5)+AddOmitType(3, 4))

	a, b := Swap("abc", "bcd")
	fmt.Println("Multiple return values.", a, b)

	year, month, day := SplitDate("1991-11-16")
	fmt.Printf("Year: %d\nMonth: %d\nDay: %d\n", year, month, day)
}
