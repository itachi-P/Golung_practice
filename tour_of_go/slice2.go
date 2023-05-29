package main

import "fmt"

//Creating a slice with make
//組み込み関数makeによって動的サイズの配列も作成可能
//ゼロ化された配列を割り当て、その配列を指すスライスを返す

func main() {
	//中の要素5つが全て0で初期化された配列を指すSliceを返す
	a := make([]int, 5)
	fmt.Println(a)
	//makeの3番目の引数にスライスの容量(capacity)を指定可能
	b := make([]int, 0, 5) //中は空で容量が5の配列を指すSlice
	fmt.Println(b, "capacity:", cap(b))

	b = b[1:cap(b)] //0番目の要素を切り落とす
	fmt.Println(b, "capacity:", cap(b))

	a = []int{4, 5, 6, 7, 8}
	b[0], b[1], b[2] = 1, 2, 3
	fmt.Println(a, cap(a), b, len(b), cap(b))

	//capacityを超えて配列長を伸ばすこともできる
	//内部的にはこの挙動は別のSliceとして新規生成される
	b = append(b, a...) //aをまるごとbの末尾に追加
	//capacity<=5までは元の配列またはスライスへの参照
	//capacityが参照元の容量を超えた時点で別物として生成
	fmt.Println(b, len(b), "capacity:", cap(b))
	//※新規に割り当てられる容量が合計要素数と等しいとは限らない
}
