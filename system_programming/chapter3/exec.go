package chapter3

import (
	"io"
	"os"
)

func Run() {
	{
		// Q3.1
		src, err := os.Open("./chapter3/q3_1.txt")
		if err != nil {
			panic(err)
		}
		defer src.Close()
		dst, err := os.Create("./chapter3/q3_2.txt")
		if err != nil {
			panic(err)
		}
		defer dst.Close()
		io.Copy(dst, src)
	}
}
func DryRun() {
}
