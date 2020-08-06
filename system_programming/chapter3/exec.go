package chapter3

import (
	"crypto/rand"
	"io"
	"os"
)

func Run() {
	if false {
		{
			// Q3.1
			src, err := os.Open("./chapter3/q3_1_1.txt")
			if err != nil {
				panic(err)
			}
			defer src.Close()
			dst, err := os.Create("./chapter3/q3_1_2.txt")
			if err != nil {
				panic(err)
			}
			defer dst.Close()
			io.Copy(dst, src)
		}
	} // false end
	{
		// Q3.2
		f, err := os.Create("./chapter3/q3_2.txt")
		if err != nil {
			panic(err)
		}
		buf := make([]byte, 1024)
		_, err = rand.Read(buf)
		if err != nil {
			panic(err)
		}
		f.WriteString(string(buf))
	}
}
func DryRun() {
}
