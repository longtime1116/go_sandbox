package main

import (
	"fmt"
	"math"
	"time"
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

func Announce(m string) {
	// 関数リテラルをうまく使うと良い
	go func() {
		time.Sleep(time.Second)
		fmt.Println(m)
	}()
}

func main() {
	if false {
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

		Announce("announced!")
		time.Sleep(2000 * time.Millisecond)
	}
	{
		// チャンネルを使った便利なイディオム！
		c := make(chan int)
		go func() {
			// 時間のかかる処理
			time.Sleep(time.Second)
			// なんでもいいから値を入れる
			c <- 1
		}()
		fmt.Println("waiting...")
		// 値を受け取るまで待ってくれる
		<-c
		fmt.Println("done!")
	}
	{
		// 第二引数でバッファする数を設定できるので、これをセマフォとして活用できる
		semaphore := make(chan int, 5)
		for i := 0; i < 10; i++ {
			go func(i int) {
				semaphore <- 1
				fmt.Println("i:", i)
			}(i)
		}
		fmt.Println("waiting...")
		time.Sleep(time.Second * 2)
		for i := 0; i < 5; i++ {
			<-semaphore
		}
		time.Sleep(time.Second)
		// このforでchannelから値を取り出して初めてiのPrintlnが実行される
		for i := 0; i < 5; i++ {
			<-semaphore
		}
	}
}
