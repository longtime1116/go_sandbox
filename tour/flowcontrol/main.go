package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
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

// defer に渡した関数の評価は直ちにされるが、実行は呼び出し元の関数がreturn後になされる
// NOTE: file の close とか、postprocess につかえそう
func execDefer1() {
	x := 0
	defer fmt.Println("x:", x)
	x++
	fmt.Println("x:", x)
}

// LIFOの順番で実行される
func execDefer2() {
	fmt.Println("counting")
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
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

	// switchはcaseの末尾に自動的にbreakステートメントが埋め込まれている！
	// 定数でなくても良いし、整数でなくても良い
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When is Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday {
	case today:
		fmt.Println("Today.")
	case (today + 1) % 7:
		fmt.Println("Tomorrow.")
	case (today + 2) % 7:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	execDefer1()
	execDefer2()
}
