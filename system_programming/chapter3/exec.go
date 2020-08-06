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

func CopyN(dst io.Writer, src io.Reader, n int64) (written int64, err error) {
	lr := io.LimitReader(src, n)
	return io.Copy(dst, lr)
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
	} // false end
	{
		// Q3.6
		var (
			computer    = strings.NewReader("COMPUTER")
			system      = strings.NewReader("SYSTEM")
			programming = strings.NewReader("PROGRAMMING")
		)
		var stream io.Reader
		r_a := io.NewSectionReader(programming, 5, 1)
		r_s := io.LimitReader(system, 1)
		r_c := io.LimitReader(computer, 1)
		r_i := io.NewSectionReader(programming, 8, 1)
		// not use pipe
		//r_i2 := io.NewSectionReader(programming, 8, 1)
		//stream = io.MultiReader(r_a, r_s, r_c, r_i, r_i2)
		// use pipe
		pr, pw := io.Pipe()
		writer := io.MultiWriter(pw, pw)
		// QUESTION: io.Pipe のブロッキングを回避するために、io.CopyNを使っているらしいが、意味がよくわからない。
		// 			 LimitReaderの方が字数指定していることに意味があるのはわかるが。。
		// go io.CopyN(writer, r_i, 1)
		go io.Copy(writer, r_i)
		defer pw.Close()
		stream = io.MultiReader(r_a, r_s, r_c, io.LimitReader(pr, 2))

		io.Copy(os.Stdout, stream)
	}
}

func DryRun() {
}
