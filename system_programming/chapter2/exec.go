package chapter2

import (
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
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

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	source := map[string]string{
		"Hello": "World",
	}

	zw := gzip.NewWriter(w)
	mw := io.MultiWriter(zw, os.Stdout)
	encoder := json.NewEncoder(mw)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(source); err != nil {
		log.Print(err)
	}
	if err := zw.Close(); err != nil {
		log.Print(err)
	}
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
		{
			// Q2.2
			wStdout := csv.NewWriter(os.Stdout)
			s := []string{"aaa", "bbb"}
			wStdout.Write(s)
			wStdout.Flush()
		}
	}
	{
		// Q2.3
		http.HandleFunc("/", handler)
		http.ListenAndServe(":8080", nil)

	}
}

func DryRun() {
}
