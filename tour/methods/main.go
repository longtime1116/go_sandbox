package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
	"os"
	"strings"
	"time"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/reader"
)

type Vertex struct {
	X, Y float64
}
type MyFloat float64

type Abser interface {
	Abs() float64
}

type I interface {
	M()
}

type T struct {
	S string
}
type F float64

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func (f F) M() {
	fmt.Println(f)
}

// クラスは無いが、型にメソッドを定義できる
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func Scale(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

func (mf MyFloat) Abs() (ret float64) {
	if mf < 0 {
		ret = float64(-mf)
	} else {
		ret = float64(mf)
	}
	return
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("int: %T, %v\n", v, v)
	case string:
		fmt.Printf("string: %T, %v\n", v, v)
	default:
		fmt.Printf("I dont know about type %T\n", v)
	}

}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("<name: %v, age: %v>", p.Name, p.Age)
}

type IPAddr struct {
	addr1, addr2, addr3, addr4 int
}

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip.addr1, ip.addr2, ip.addr3, ip.addr4)
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s(%v)", e.What, e.When)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work.",
	}
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0

	c := 0
	for d := 1.0; math.Abs(d) > 1e-10; z -= d {
		d = (z*z - x) / (2 * z)
		c++
	}
	fmt.Println("count:", c)

	return z, nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// %v を使うとstack overflowする
	// これは、ErrNegativeSqrtを文字列に変換する String() のなかで、また ErrNegativeSqrt を文字列にするために String()を読んで・・・となるから
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

type MyReader struct{}

func (r MyReader) Read(buf []byte) (n int, e error) {
	for i := range buf {
		buf[i] = 'A'
	}
	return len(buf), nil
}

type rot13Reader struct {
	r io.Reader
}

func (r3r rot13Reader) Read(buf []byte) (int, error) {
	n, err := r3r.r.Read(buf)
	if err == io.EOF {
		return n, err
	}
	for i, c := range buf {
		if (c >= 'A' && c <= 'M') || (c >= 'a' && c <= 'm') {
			buf[i] = c + 13
		} else if c >= 'N' && c <= 'Z' || c >= 'n' && c <= 'z' {
			buf[i] = c - 13
		}
	}
	return n, nil
}

type Image struct {
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 1000, 2000)
}

func (i Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 0, 255}
}

func main() {
	if false {
		{
			// NOTE: ポインタレシーバを使う理由
			// 1. メソッドがレシーバの指す先の変数を変更するため
			// 2. メソッドの呼び出しごとに変数のコピーをせずに済む
			// NOTE: 一般に、変数レシーバとポインタレシーバを混合させるべきでは無い

			v := Vertex{3, 4}
			fmt.Println(v.Abs())
			fmt.Println(Abs(v))

			//fmt.Println(-math.Sqrt2)
			mf := MyFloat(-math.Sqrt2)
			fmt.Println(mf.Abs())

			// (&v).Scale(10) を自動で呼び出してくれる
			// また、変数レシーバのメソッドにポインタを渡したら(*p).Abs()みたいに呼び出してくれる
			v.Scale(10)
			(&v).Scale(10)
			// ポインタレシーバじゃなければ、明示的にポインタを渡す必要あり
			Scale(&v, 10)
			fmt.Println(v)
		}
		{
			// 普通は mf や v に Abs() を呼び出すが、aに対して呼び出せる
			var a Abser
			mf := MyFloat(-math.Sqrt2)
			v := Vertex{3, 4}

			a = mf // a Myfloat implements Abser
			a = &v
			//a = v // compile error
			fmt.Println(a.Abs())

			// インタフェースを実装することを明示的に宣言する必要はない
			// この場合、TはIというinterfaceを実装することを明示的に宣言しない
			var i I = &T{"hello"}
			describe(i)
			i.M()
			i = F(math.Pi)
			describe(i)
			i.M()
		}
		{
			var i I
			var t *T
			i = t
			// nil に対して実行をしてもSEGVにならないように実装するのが一般的
			// nilを保持するインタフェースそれ自体はnilではない
			i.M()
			describe(i)
			// NOTE: tは別にインスタンスではないので、あくまで引数としてT型が渡されたときのM()関数が呼ばれるから、こういう挙動になる
			t.M()
		}
		{
			var i I
			describe(i)
			// runtime error(SIGSEGV)
			//i.M()
		}
		{
			// 空のインターフェース
			var i interface{}
			describe(i)
			i = 42
			describe(i)
		}
		{
			var i interface{} = "hello"
			s := i.(string)
			fmt.Println(s)
			s, ok := i.(string)
			fmt.Println(s, ok)
			// 型アサーションによりpanicを引き起こす
			//f := i.(float64)
			//fmt.Println(f)
			f, ok := i.(float64)
			fmt.Println(f, ok)
			var i2 I = &T{"hello"}
			t, ok := i2.(*T)
			fmt.Println(t, ok)
		}
		{
			do(21)
			do("hello")
			do(true)
		}
		{
			// Person型
			a := Person{"Adam", 28}
			b := Person{"Bob", 77}
			fmt.Println(a)
			fmt.Println(b)
		}
		{
			// Exercise: IPAddr 型
			ip := IPAddr{192, 168, 11, 1}
			fmt.Println(ip)
			hosts := map[string]IPAddr{
				"lookback":  {127, 0, 0, 1},
				"googleDNS": {8, 8, 8, 8},
			}
			for name, ip := range hosts {
				fmt.Printf("%v: %v\n", name, ip)
			}
		}
		{
			if err := run(); err != nil {
				fmt.Println(err)
			}
		}
		{
			// Exercise: Errors
			x, err := Sqrt(math.Sqrt2)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(x)
			}
			y, err := Sqrt(-math.Sqrt2)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(y)
			}

		}
		{
			// Reader
			r := strings.NewReader("Hello, Reader!")
			b := make([]byte, 8)
			for {
				// io.Reader インタフェースは↓を持つ
				// 	func (T) Read(b []byte) (n int, err error)
				n, err := r.Read(b)
				fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
				fmt.Printf("b[:n] = %q\n", b[:n])
				if err == io.EOF {
					break
				}
			}
		}
		{
			// Exercise: Readers
			reader.Validate(MyReader{})
			//r := MyReader{}
			//b := make([]byte, 4)
			//for {
			//	// io.Reader インタフェースは↓を持つ
			//	// 	func (T) Read(b []byte) (n int, err error)
			//	n, err := r.Read(b)
			//	fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
			//	fmt.Printf("b[:n] = %q\n", b[:n])
			//	if err == io.EOF {
			//		break
			//	}
			//}
		}
		{
			// Exercise: rot13Reader
			s := strings.NewReader("Lbh penpxrq gur pbqr!")
			r := rot13Reader{s}
			io.Copy(os.Stdout, &r)
		}
	} // false end

	{
		// Exercise: Image
		// 	./main |  sed -e 's/IMAGE:\(.*\)/<img src="data:image\/png;base64,\1">/g' > hoge.html; open hoge.html
		m := Image{}
		pic.ShowImage(m)
	}
}
