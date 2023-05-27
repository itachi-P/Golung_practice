package main

import "fmt"

//pointer:値のメモリ上のアドレスを指す
/* 基本的にC言語と同じだが、ポインタ演算は無い */

func main() {
	var p *int     //ポインタの宣言（初期化なし）
	fmt.Println(p) //初期値はnilになる

	i, j := 256, 273

	//&オペレータはそのオペランドへのポインタを引き出す
	p = &i //int型のポインタpにiの値の入ったアドレスを渡す

	//ポインタそのものを見てみる
	fmt.Printf("Type:%T Value:%p\n", p, p)
	//ポインタの指す実体値(参照先)を読み出す
	fmt.Printf("Type:%T Value:%d\n", *p, *p)
	*p = 2147483648      //ポインタを通してiへ値を代入する
	fmt.Println("i:", i) //変数iを直接触らずしてiの値が変更される

	p = &j               //ポインタpの参照先をiからjに書き換え
	*p = *p / 3          //ポインタpを通してjに対し除算を行う
	fmt.Println("j:", j) // see the new value of j

	//ポインタの指すアドレスもjのものに変わっている
	fmt.Printf("Type:%T Value:%p\n", p, p)
	fmt.Printf("&i:%p &j:%p\n", &i, &j)
}
