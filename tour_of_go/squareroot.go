package main

import (
	"fmt"
	"math"
)

//ニュートン法を用いて平方根の近似値を求める

func Sqrt(x float64) float64 {
	z := float64(1) //2乗して入力値xに近付ける
	var s float64   //z^2がxにどれくらい近付いたかの比較用

	//最大でも試行10回までとする
	for i := 1; i <= 10; i++ {

		//ニュートン法（とりあえず理解できなくてもOKとする）
		z -= (z*z - x) / (2 * z)

		//%fだと精度(表示桁数)が足りないので%g（計算結果には影響しない）
		fmt.Printf("%d回目: %g 差分: %g\n", i, z, math.Abs(z-s))

		//前回の値との差分が0.0000000001を下回ったら中断し最終的な解とする
		if math.Abs(z-s) < 1e-10 { //1e-1=0.1,1e-10なら0が10個

			/* 上記if文の別バージョン
			本家のmath.Sqrt()と同じ値になったら抜ける
			if math.Sqrt(x) == z { */
			break
		}

		//1つ前の値との差分を比較する
		s = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println("一夜一夜に人見頃")

	fmt.Println(Sqrt(3))
	fmt.Println("人並みに奢れや")
}
