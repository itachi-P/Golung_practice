package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 標準入力から1行取得するためのスキャナーを作成
	sc := bufio.NewScanner(os.Stdin)
	// 1行目の入力値(2つの整数値の繰り返し入力回数)を取得
	sc.Scan()
	// 文字列を数値に変換
	var n, _ = strconv.Atoi(sc.Text())

	// 2行目以降の入力値を1行目で入力された回数分取得し、指定のフォーマットに従って出力
	for i := 0; i < n; i++ {
		sc.Scan()
		var s = strings.Split(sc.Text(), " ")
		fmt.Println("hello = " + s[0] + " , world = " + s[1])
	}
}
