package main

import "fmt"

//クロージャを用いたフィボナッチ数の実装演習

func fibonacci() func(int) int {
	x, y, tmp := 0, 1, 0
	return func(i int) int {
		//0番目と1番目の要素は0と1で固定
		if i < 2 {
			return i
		}
		//フィボナッチ数の次の項(前の2つの要素の和)の計算
		tmp = x + y
		x = y
		y = tmp
		//xとyの値を更新保持しつつ、戻り値としては古いxとyの和を返す
		return tmp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		//元の問題では引数無しだったが、引数にiを持たせた
		fmt.Println(f(i))
	}
}
