package main

import "fmt"

//クロージャ
//Goの関数はclosure:それ自身の外部から変数を参照する関数値

func adder() func(int) int {
	sum := 0
	return func(x int) int { //このfunc(int)がクロージャ
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
