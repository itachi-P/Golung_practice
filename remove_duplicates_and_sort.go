package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//集合(スライス)の結合とソート
//数列AまたはBに含まれる値をすべて列挙し、重複を取り除き昇順に出力
//Goには重複を取り除く標準関数がない。
//その為、一般的にmap[string]bool(かint)を使って自作する必要がある

/*
並べ替えなど、数列の大小比較を文字列のまま行うとASCIIコードの総和の比較になる
ASCIIコードでの比較は必ずしも数値の大小と一致しない為、数値に変換する必要がある
たとえば"10" < "2"など、桁数の少ない方が大きくなったりしてしまう
*/

func main() {
	sc := bufio.NewScanner(os.Stdin)
	//数列A, Bの要素をint型のスライスに格納
	sc.Scan()
	//int型に変換して格納(大小比較や計算の手前までは文字列のままでも良い)
	var a []int //int型のスライスを作成
	for _, v := range strings.Split(sc.Text(), " ") {
		n, _ := strconv.Atoi(v)
		a = append(a, n)
	}
	sc.Scan()
	var b []int
	for _, v := range strings.Split(sc.Text(), " ") {
		n, _ := strconv.Atoi(v)
		b = append(b, n)
	}

	//数列AとBを合わせたスライスを作成し、重複を取り除く
	merged := append(a, b...) //...は可変長引数を展開する
	c := removeDuplicates(merged)

	//重複を取り除いたスライスを昇順にソート
	//この手前までは文字列のままでも良いが、並べ替えや計算は数値に変換する必要がある
	sort.Ints(c)
	for i := 0; i < len(c); i++ {
		fmt.Print(c[i])
		if i < len(c)-1 {
			fmt.Print(" ")
		}
	}
}

// Goには重複を取り除く関数がないので自作
func removeDuplicates(elements []int) []int {
	// 重複を取り除くためのmapを作成
	seen := make(map[int]bool) //boolでなくint=0->1のmapでも可
	uniques := []int{}

	for _, element := range elements {
		// まだ一度も出現していない要素ならuniquesに追加
		if !seen[element] {
			uniques = append(uniques, element)
			seen[element] = true
		}
	}
	return uniques
}
