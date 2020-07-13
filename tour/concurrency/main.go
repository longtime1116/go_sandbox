package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int, d time.Duration) {
	time.Sleep(d)
	fmt.Println("sum called", s)
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c

}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// ここではこれをしないとデッドロックを検出してエラーになる
	// とはいえ、チャネルはファイルと違って必ずcloseしなければならないわけではない
	close(c)
}

func main() {
	if false {
		{
			// goroutine(Goのランタイムに管理される軽量なスレッド)
			go say("world")
			say("hello")
		}
		{
			s := []int{7, 2, 8, -9, 4, 0}
			// Channel 型の変数を作成(intの送受信用)
			c := make(chan int)
			go sum(s[:len(s)/2], c, 1000*time.Millisecond)
			// 当然↑よりも↓の方が、500ms早く実行される
			go sum(s[len(s)/2:], c, 500*time.Millisecond)
			time.Sleep(2000 * time.Millisecond)
			fmt.Println("start receiving")
			// 500msの方が早いのでそちらが最初に受信されてxに入る
			x, y := <-c, <-c // receive from c
			fmt.Println(x, y, x+y)
			//x := <-c
			//fmt.Println(x)
			//y := <-c
			//fmt.Println(y)
		}
		{
			ch := make(chan int, 2)
			ch <- 13
			ch <- 29
			// ↓バッファの長さを超えるのでエラー
			// ch <- 31
			fmt.Println(<-ch)
			ch <- 31
			fmt.Println(<-ch)
			v, ok := <-ch
			fmt.Println(v, ok)
			close(ch)
			v, ok = <-ch
			fmt.Println(v, ok)

		}
		{
			c := make(chan int, 10)
			go fibonacci(cap(c), c)
			// チャネルが閉じられるまで値を受信し続ける
			for i := range c {
				fmt.Println(i)
			}
		}
		{
			c := make(chan int)
			quit := make(chan int)
			go func() {
				for i := 0; i < 10; i++ {
					fmt.Println(<-c)
				}
				quit <- 0
			}()
			fibonacci2(c, quit)
		}
	} // false end

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	// このSleepを入れると、selectで複数のcaseに当てはまる時にランダムで選ばれることを実証できる
	// time.Sleep(1000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("   .")
			time.Sleep(50 * time.Millisecond)
		}
	}

}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		fmt.Println("waiting...")
		select {
		case c <- x:
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("sent %v to c\n", x)
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
