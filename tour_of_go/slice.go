package main

import "fmt"

//可変長のSliceと固定長配列Array
/* Goでは配列(Array)は固定長
スライスは配列（の部分列）への参照のようなもの
実はスライス自体にはどんなデータも格納していない */

func main() {
	//[6]と明記する代わりに[...]で自動的に要素数が計測される
	arr_primes := [...]int{2, 3, 5, 7, 11, 13}
	//元の配列の1~3番目の要素を切り出し新たなSliceを生成
	var s []int = arr_primes[1:4] //0及び4番目は含まれない
	fmt.Println(s)                //=>{3, 5, 7}

	//配列から切り出したSliceを変更すると参照元の対応する要素も変更される
	s[0] = 100
	fmt.Println(s, arr_primes)

	//可変長であるSliceには任意の数の要素を追加可能
	s = append(s, arr_primes[4], arr_primes[5])
	s[2] = 200
	fmt.Println(s, arr_primes)

	//元の配列との対応部分を超えて要素が追加されると、以降は参照元が別物になる
	s = append(s, 17) //要素数は同数の6でも参照元のindexを超える
	s[3] = 300
	fmt.Println(s, arr_primes) //Sliceの方だけが変更されるようになる

	//別のSliceを定義
	prime2 := []int{19, 23, 29, 31, 37}
	//[n](または[...])だとArray,[]はSliceで別の型になる
	fmt.Printf("arr_primesの型:%T, prime2の型:%T\n", arr_primes, prime2)
	prime2 = append(prime2, 41, 43, 47)

	//任意の数の要素でなく、別のSliceも追加可能
	s = append(s, prime2...) //...は要素番号(index)の省略記法、つまり全体を表す
	fmt.Println(s)

	//文字列も同様
	//sliceは長さ(length)と容量(capacity)の要素を持つ
	//参照元に変更が反映されなくなるのはこの容量を超えた時

	//固定長配列を要素数を明記せず生成
	array := [...]string{"PHP", "Java", "Ruby"}
	//固定長配列の部分列(start : end-1)からのスライス作成
	//[:]と省略した場合は先頭から最後尾、つまり全要素
	slice := array[1:]
	fmt.Println(len(slice)) //要素数は"Java","Ruby"の2つ
	//スライスの最初の要素から数えて、元となる配列に対応するindex
	fmt.Println(cap(slice), cap(array)) //cap(array)は固定値3

	slice[0] = "Golang"       //slice[0]の要素を変更
	fmt.Println(slice, array) //元の配列の対応する要素も変更される

	//appendによって元の配列の対応indexを超えた時点で参照元が変わる
	//"Shakespeare"は元の配列には存在しない4番目の要素、index3番にあたる
	slice = append(slice, "Shakespeare")
	fmt.Println(cap(slice), cap(array)) //元の配列arrayの容量を超える

	slice[1] = "Haskell"
	fmt.Println(slice, array) //元の配列の参照が外れた為、変更されない
}
