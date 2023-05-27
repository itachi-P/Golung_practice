package main

import "fmt"

//pointer:値のメモリアドレスを指す
/*
基本的にC言語と同じだが、ポインタ演算は無い

ポインタの宣言(初期値はnilになる)
var p *int
&オペレータはそのオペランドへのポインタを引き出す
i := 16
p = &i
fmt.Println(*p) //ポインタの指す実体値を読み出す
*p = 256 		//ポインタを通してiへ値を代入する
*/

func main() {
	var p *int //ポインタの宣言（初期化なし）
	fmt.Println(p)

	i, j := 42, 2701

	p = &i          // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
