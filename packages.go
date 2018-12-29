package main

import (
	"fmt"
)

func add_omit_type(x, y int) int {
	return x + y
}

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("My favorite number is", add(2, 5)+add_omit_type(3, 4))
}
