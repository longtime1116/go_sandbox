package main

import (
	"fmt"
	"strconv"
	"strings"
)

var package_variable1 bool                                  // package 内で利用可能
var package_variable2, package_variable3 bool = true, false // 初期化
var package_variable4, package_variable5 = true, false      // 初期化子があれば型省略可能

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

	var c, python, java = true, false, "No!"
	fmt.Println("variables")
	fmt.Println(c, python, java)
	fmt.Println(package_variable1, package_variable2, package_variable3, package_variable4, package_variable5)

	i := 0 // var 宣言を使わず暗黙的に型宣言
}
