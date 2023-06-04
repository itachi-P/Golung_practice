package main

import "fmt"

//クロージャ
//Goの関数はclosure:それ自身の外部から変数を参照する関数値

func adder() func(int) int {
	//各クロージャがこの内部変数を個別に持つ（同じ参照元ではない）
	sum := 0
	return func(x int) int { //この無名関数func(int)がクロージャ
		sum += x
		return sum
	}
}

func main() {
	//元の内部無名関数にそれぞれpos,negという名前を付け生成
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
	fmt.Printf("Type:%T, Address:%p\n", pos, &pos)
	fmt.Printf("Type:%T, Address:%p\n", neg, &neg)
	//クロージャの内部変数に直接アクセスはできない
	//fmt.Println(pos.sum)
}
