package main

import (
	"fmt"
	"math"
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

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
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
}
