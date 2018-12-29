package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("My favorite number is", add(2, 5))
}
