package main

import (
	"fmt"
	"math"
)

//function value(関数値)
//関数も変数として扱える。引数や戻り値にもできる。

// ややこしいが2つのfloat64を引数に取りfloat64を返す内部関数を引数にしている
// ここでは内部関数側では実装ロジックを書かず、main側で定義している
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	//上で定義した(x^2+y^2)の平方根に3, 4が渡され実行される
	fmt.Println(compute(hypot))
	//引数に組み込み関数であるmath.Powを渡しているので3^4が実行される
	fmt.Println(compute(math.Pow))
}
