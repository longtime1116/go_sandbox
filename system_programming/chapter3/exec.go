package chapter3

import (
	"archive/zip"
	"crypto/rand"
	"io"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=sample.zip")

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()
	w1, err := zipWriter.Create("./file1.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(w1, strings.NewReader("Hello, world! This is file1.\n"))
	w2, err := zipWriter.Create("./file2.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(w2, strings.NewReader("Hello, world! This is file2.\n"))
}

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
		{
			// Q3.2
			// 1
			f, err := os.Create("./chapter3/q3_2_1")
			if err != nil {
				panic(err)
			}
			defer f.Close()
			buf := make([]byte, 1024)
			_, err = rand.Read(buf)
			if err != nil {
				panic(err)
			}
			f.WriteString(string(buf))
			// 2
			f, err = os.Create("./chapter3/q3_2_2")
			if err != nil {
				panic(err)
			}
			defer f.Close()
			_, err = io.CopyN(f, rand.Reader, 1024)
			if err != nil {
				panic(err)
			}
		}
		{
			// Q3.3
			// 1. q3_3_dest.gz を unzip すると q3.txt が展開され、中にはHello, Worldが書かれている
			f, err := os.Create("./chapter3/q3_3_dst.zip")
			if err != nil {
				panic(err)
			}
			defer f.Close()
			zipWriter := zip.NewWriter(f)
			defer zipWriter.Close()
			w1, err := zipWriter.Create("./file1.txt")
			if err != nil {
				panic(err)
			}
			io.Copy(w1, strings.NewReader("Hello, world! This is file1.\n"))
			w2, err := zipWriter.Create("./file2.txt")
			if err != nil {
				panic(err)
			}
			io.Copy(w2, strings.NewReader("Hello, world! This is file2.\n"))
		}
		{
			// Q3.4
			http.HandleFunc("/", handler)
			http.ListenAndServe(":8080", nil)
		}
	} // false end
	{
		// Q3.5
		f, err := os.Open("./chapter3/q3_5.txt")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		_, err = CopyN(os.Stdout, f, 32)
		if err != nil {
			panic(err)
		}
		os.Stdout.Write([]byte("\n"))
	}
}

func CopyN(dst io.Writer, src io.Reader, n int64) (written int64, err error) {
	lr := io.LimitReader(src, n)
	return io.Copy(dst, lr)
}

func DryRun() {
}
