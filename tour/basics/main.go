package main

import (
	"fmt"
	"math"

	// factored インポートステートメント。こっちの方が望ましい
	"math/cmplx"
	"math/rand"
)

// TODO: Go's Declaration Syntax
// 		 https://blog.golang.org/declaration-syntax
func add1(x int, y int) int {
	return x + y
}

// 型が同じ場合は省略可能
func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// 短い関数では戻り値となる変数に名前を最初につけてしまう naked return value が便利
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 変数宣言
var c, python, java bool

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println("number: ", math.Sqrt(10))
	fmt.Println("random: ", rand.Intn(10))
	// 大文字で始まるものは、外部パッケージから参照できるエクスポートされたもの
	fmt.Println("pi: ", math.Pi)

	fmt.Println(add1(1, 2))
	fmt.Println(add2(3, 4))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(19))

	var i int
	fmt.Println(i, c, python, java)

	var p, q int = 1, 2
	// 暗黙的な型宣言は:=で
	s := 3
	c, python, java := true, false, "no!"
	fmt.Println(p, q, s, c, python, java)

	// 型
	// NOTE: 理由がないならintを使う
	var (
		ToBe      bool       = true
		MaxUint              = ^uint(0)
		MaxUint64 uint64     = 1<<64 - 1
		z         complex128 = cmplx.Sqrt(-5 + 12i)
		str       string     = "mr.hoge"
	)
	fmt.Printf("Types: %T, Value: %v\n", ToBe, ToBe)
	fmt.Printf("Types: %T, Value: %v\n", MaxUint, MaxUint)
	fmt.Printf("Types: %T, Value: %v\n", MaxUint64, MaxUint64)
	fmt.Printf("Types: %T, Value: %v\n", z, z)
	fmt.Printf("Types: %T, Value: %v\n", str, str)

	// type conversions
	var (
		z1 int     = 2
		f  float64 = math.Sqrt(float64(z1))
		z3         = uint(f)
	)
	fmt.Println(z1, f, z3)

	// constants は := で宣言できない
	const Word = "tango"
	fmt.Println("Hello,", Word)

	const (
		Big   = 1 << 100
		Small = Big >> 99
	)
	// エラーになる
	//fmt.Println(Big, Small)
	fmt.Println(needInt(Small), needFloat(Small), needFloat(Big))
}
