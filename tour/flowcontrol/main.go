package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// Sqrt 大文字で始まる関数はexportされるのでコメントが必要
func Sqrt(x float64) float64 {
	z := 1.0

	c := 0
	for d := 1.0; math.Abs(d) > 1e-10; z -= d {
		d = (z*z - x) / (2 * z)
		c++
	}
	fmt.Println("count:", c)

	return z
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	// for の初期化と後処理は省略できるし、;も省略できるので、whileが不要となる
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)

	// infinite loop
	//for {
	//}

	fmt.Println(sqrt(-25))

	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
	fmt.Println(Sqrt(25))

}
