package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	hundleName := removeVowels(scanner.Text())
	fmt.Println(hundleName)
}

// 母音を除外した文字列を生成する関数
func removeVowels(input string) string {
	vowels := "aeiouAEIOU"

	//strings.Map()は文字列の各文字に対し、指定した無名(匿名)関数を1文字づつ実行する
	//func Map(mapping func(rune) rune, s string) string {
	//第1引数には文字列の各文字に対し実行する匿名関数、第2引数に対象の文字列を指定する
	//Mapは内部関数mappingに従い、文字列sのすべての文字を変更したコピーを返す。
	//なんなら文字列の長さを含め、全く原型を留めない文字列を返すことも可能。
	//マッピングが負の値を返した場合、その文字は文字列から削除され置換されない。
	output := strings.Map(func(r rune) rune {
		//ContainsRuneは文字列sにルーンrが含まれているかどうかの真偽値を返す
		//runeはint32の別名で、Unicodeポイント('0'=48,'A'=65,'a'=97など)を表す
		if strings.ContainsRune(vowels, r) {
			return -1 // 母音を除外するために-1を返す
		}
		return r // 母音以外の文字は置換せずそのまま返す
		//return 'a' //左記の場合、母音以外は全て'a(97)'に置換される: "example"(->"xmpl")->"aaaa"
	}, input)

	return output

}
