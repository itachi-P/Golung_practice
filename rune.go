package main

import (
	"bufio"
	"fmt"
	"os"
)

// Golangにおける文字列はrune(ルーン文字)のスライスである
// runeは文字列を構成する文字の型(Unicodeのコードポイント)である
// 1文字は1~4バイトで表現される(日本語や絵文字は3~4バイト)
// runeはint32の別名型であり、0~127の範囲の値はASCII文字に対応している
// runeリテラルはシングルクォートで囲み、verbは%U, %c, %q, %d等が使える
// 普通に表示するだけなら%vや%cでも良い(%s不可)が、%qを使った方が可読性が高い
// %UはUnicode表記、%qはクォート付き文字、%dは10進数表記

func main() {
	r1 := 'a'    //アルファベットや数字や特殊記号はASCII文字(0~127)に対応
	rune2 := 'あ' //10進数表記(%d)で12354、16進数表記(Unicode)で3042
	rune3 := '🍣'
	fmt.Printf("%U %q %d\n", r1, r1, r1)          //U+0061, 'a', 97
	fmt.Printf("%U %q %d\n", rune2, rune2, rune2) //U+3042, 12354
	fmt.Printf("%U %q %d\n", rune3, rune3, rune3) //U+1F363, 127843

	//標準入力から文字列を受け取り、1文字ずつrune型で出力する
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("何か文字列を入力してください。(特殊記号やマルチバイト文字も可))")
	sc.Scan()
	str := sc.Text()
	//マルチバイト文字の場合indexは+1文字分ではなく、バイト数分進む
	fmt.Println("index\tchar\t10進数\tUnicode(16進数)")
	for i, r := range str {
		fmt.Printf("%d\t%q\t%d\t%U\n", i, r, r, r)
	}
}
