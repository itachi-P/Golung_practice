package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

/* A Tour of Go : Exercise 03(Maps)
https://go-tour-jp.appspot.com/moretypes/23
WordCount関数の実装
string s で渡される文章の、各単語の出現回数のmapを返す
（スペース区切りの英文であることが前提）
用意されているInput（一例）:
"I ate a donut. Then I ate another donut."
上記のInputに対し、戻り値を下記のようなMapに加工する:
map[string]int{"I": 2, "ate": 2, "a": 1, "donut.": 2, "Then": 1, "another": 1,}
*/

func WordCount(s string) map[string]int {
	//公式が用意してある英文をまずはスペースで区切りスライスに分解
	words := strings.Fields(s)
	fmt.Printf("%q\n", words)
	//戻り値として返すMapを準備
	m := make(map[string]int)

	for _, v := range words {
		//同じ単語の出現回数（基本は1）
		count := 0
		//もし既にKeyが存在すれば2以上にカウントアップする
		if v, ok := m[v]; ok { //2つ目は変数ok == trueの省略記法
			count = v
		}
		//MapにKey（単語）とValue（出現回数）を追加または上書きする
		m[v] = count + 1
		//fmt.Println(v)
	}
	return m
}

func main() {
	/*問題を解く上で、この外部パッケージ関数の中身の理解が必要
	https://github.com/golang/tour/blob/master/wc/wc.go
	Test関数によりWordCount関数が期待される挙動と一致するか判定される
	引数(s string)となる英文は自動的に与えられる
	*/
	wc.Test(WordCount)
}
