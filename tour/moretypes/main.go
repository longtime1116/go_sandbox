package main

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

// Vertex 構造体は、フィールド(field)のあつまり
type Vertex struct {
	X int
	Y int
}

type Location struct {
	Lat, Long float64
}

func printSlices(s []int) {
	// スライスの容量は、スライスの最初の要素から数えて、元となる配列の要素数です。
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printBoard(b [][]string) {
	for i := 0; i < len(b); i++ {
		fmt.Printf("%s\n", strings.Join(b[i], " "))
	}
}

// Pic is for exercise
func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		p[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			p[i][j] = uint8(i + j)
		}
	}

	return p

}

// WordCount is for exercise
func WordCount(s string) map[string]int {
	m := make(map[string]int)

	for _, v := range strings.Fields(s) {
		m[v]++
	}

	return m
}

func compute(fn func(float64, float64) float64, x, y float64) float64 {
	return fn(x, y)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	pp := -1
	p := 1
	return func() int {
		cur := pp + p
		pp = p
		p = cur
		return cur
	}
}
func fibonacci2() func() int {
	pp := 0
	p := 1
	cur := 1
	return func() int {
		pp = p
		p = cur
		cur = pp + p
		return pp
	}
}

func main() {

	// pointer
	// Cと違ってポインタ演算はないらしい！
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	// {}で作成する
	v := Vertex{1, 2}
	fmt.Println(v)
	vp := &v
	vp.X = 100
	fmt.Println(v)

	var (
		v1  Vertex = Vertex{1, 2} // Vertex 型であることは省略できるが、暗黙型というわけではない
		v2         = Vertex{X: 1} // Y: 0 is implicit
		v3         = Vertex{}     // Y: 0 is implicit
		vp1        = &Vertex{3, 4}
	)
	fmt.Println(v1, v2, v3, vp1)

	// 配列、Arrays は固定長
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// スライス、Slices は可変長。配列よりも一般的とのこと
	// 配列の参照的なもの
	var s1 []int = primes[1:4] // [1,4)
	fmt.Println("s1: ", s1)
	s2 := primes[2:]
	fmt.Println("s2: ", s2)
	s2[3] = 19
	fmt.Println("s2: ", s2)
	fmt.Println("primes: ", primes)

	// slice literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("%T\n", primes)
	fmt.Printf("%T\n", q)
	fmt.Println(q[:3])
	fmt.Println(q[3:])
	fmt.Println(q[:])

	structs := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
	}
	fmt.Println(structs)

	printSlices(q)
	printSlices(q[:0])
	printSlices(q[:4])
	printSlices(q[2:])
	printSlices(q[2:3])

	// nil slices
	var snil []int
	fmt.Println(snil, len(snil), cap(snil))
	if snil == nil {
		fmt.Println("nil!")
	}
	// creating a slice with make
	{
		a := make([]int, 5)
		b := make([]int, 0, 5)
		c := b[:2]
		d := c[2:5]
		printSlices(a)
		printSlices(b)
		printSlices(c)
		printSlices(d)
	}
	// slices of slices
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	println("init")
	printBoard(board)
	board[1][1] = "O"
	println("turn: O")
	printBoard(board)
	board[0][0] = "X"
	println("turn: X")
	printBoard(board)
	board[0][1] = "O"
	println("turn: O")
	printBoard(board)
	board[2][2] = "X"
	println("turn: X")
	printBoard(board)
	board[2][1] = "O"
	println("turn: O")
	printBoard(board)
	// appending to a slice
	{
		var s []int
		printSlices(s)
		// 容量が足りない場合はより大きいサイズの配列を割り当て直す
		s = append(s, 1)
		printSlices(s)
		s = append(s, 2, 3, 4)
		printSlices(s)
	}
	// range
	{
		pow := []int{1, 2, 4, 8, 16, 32}
		// iとvを返している
		for i, v := range pow {
			fmt.Printf("2**%d = %d\n", i, v)
		}
	}
	{
		pow := make([]int, 10)
		// vが要らないなら省略
		for i := range pow {
			pow[i] = 1 << uint(i)
		}
		// i が要らないなら_で捨てる
		for _, v := range pow {
			fmt.Printf("%d\n", v)
		}
	}
	// Exercise: Slices
	// 画像の表示方法は以下
	// 	$ ./main |  sed -e 's/IMAGE:\(.*\)/<img src="data:image\/png;base64,\1">/g' > hoge.html
	// 	$ open hoge.html
	pic.Show(Pic)

	// creating map
	{
		// ゼロ値はnil。キーを持たず、追加もできない
		var m map[string]Location
		m = make(map[string]Location)
		//var m = make(map[string]Location)
		m["Bell Labs"] = Location{40.68433, -74.39967}
		m["Alice Labs"] = Location{30.68433, -64.39967}
		fmt.Println(m["Bell Labs"])
		fmt.Println(m)

		var m2 = map[string]Location{
			"a": Location{1, 2},
			"b": {3, 4}, // 単純ならば省略可能
		}
		fmt.Println(m2)
	}
	// map操作
	{
		m := make(map[string]int)
		m["Answer"] = 42
		fmt.Println("ans:", m["Answer"])
		m["Answer"] = 48
		fmt.Println("ans:", m["Answer"])
		delete(m, "Answer")
		fmt.Println("ans:", m["Answer"])
		v, ok := m["Anwser"]
		fmt.Println("ans:", v, "Present?", ok)
	}
	wc.Test(WordCount)

	// 関数も変数
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot, 5, 12))
	fmt.Println(compute(math.Pow, 2, 5))

	// Goの関数はクロージャ(関数閉包)である
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i))
		fmt.Println(neg(-i))
	}
	// exercise: Fibonacci closure
	{
		f := fibonacci()
		for i := 0; i < 10; i++ {
			fmt.Println(f())
		}

	}
}
