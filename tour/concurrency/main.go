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

func main() {
	// goroutine(Goのランタイムに管理される軽量なスレッド)
	go say("world")
	say("hello")

}
