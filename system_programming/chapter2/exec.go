package chapter2

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
)

type Talker interface {
	Talk()
}
type Greeter struct {
	name string
}

func (g Greeter) Talk() {
	fmt.Printf("hello, I'm %v", g.name)
}

func Run() {
	if false {
		{
			// syscall の write を呼んでいるところまでデバッガで追ってみる
			fmt.Println("hoge")
		}
		{
			// インタフェース
			var talker Talker
			// インタフェースを満たすので代入できる！
			talker = Greeter{"Bob"}
			talker.Talk()
		}
		{
			// $ GOPATH=/ godoc -http ":6060" -analysis type
			// ↑を打ってから localhost:6060 で↓とかをみると、インタフェースを実装している構造体一覧が観れる！
			// http://localhost:6060/pkg/io/#Writer
		}
		{
			// Q2.1
			f, err := os.Create("q2_1.txt")
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(f, "%d, %s, %f\n", 1, "aaa", math.Pi)
			fmt.Println("q2_1.txt created!")
		}
	}
	{
		// Q2.2
		wStdout := csv.NewWriter(os.Stdout)
		s := []string{"aaa", "bbb"}
		wStdout.Write(s)
		wStdout.Flush()
	}
}

func DryRun() {
}
