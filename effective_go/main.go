package main

import (
	"fmt"
	"math"
)

type Sample struct {
	a int
	b float64
	s string
}

func f() Sample {
	return Sample{1, 2.5, "hoge"}
}

func f2(a ...interface{}) {
	fmt.Printf("%T\n", a)
	fmt.Println(a)
}

type Color int

const (
	// イオタを使ってenumを表現
	Red Color = iota
	Blue
	Yellow
)

// enum 定義したものをいい感じに出すために String() を実装しておくと良い
// QUESTION: これどうにかならんのか？
func (c Color) String() string {
	switch c {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Yellow:
		return "Yellow"
	default:
		return "Unknown"
	}
}

type ByteSize float64

// iotaを使ってこんなこともできる
const (
	_           = iota // 空の識別子に割り当てて、最初の値（０）を無視する
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	YB
)

func init() {
	// プログラム状態の正当性をチェックするのに使ったりする
	fmt.Println("init() called implicitly before main is called")
}

func main() {
	// ローカル変数のアドレスを戻り値にするのは完全に合法であり、変数領域は関数が帰った後も保持される
	s := f()
	fmt.Println(s)

	// print
	// 10進数表記、16進数表記
	var x int = 1<<4 + 1<<2
	fmt.Printf("%d %x\n", x, x)

	fmt.Printf("%T\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%#v\n", s)

	// 可変長引数
	f2(1, "abc")
	f2(1, 2, 3, 4, "abc")

	// 定数
	const Cnst = "const"
	const Pi = math.Pi
	// 式が右辺にあると、実行時にメモリ確保できないので定数とはならない。エラー。
	// const PiSqrt = math.Sqrt(math.Pi)
	fmt.Println(Cnst, Pi)
	fmt.Println(Red, Blue, Yellow)
	fmt.Println(KB, MB, GB)
}
