package main

import (
	"golang.org/x/tour/pic"
)

/* 上記import文による外部モジュールの解決の為に、
このファイル用のgo.modファイル生成、及び
ワークスペースのルートにあるgo.workにモジュールを追加したが
結局実行結果のIMAGE:の値はブラウザで表示しないと文字と記号の羅列
大人しくA Tour of Goのサンドボックス上で実行した方がよさげ */

func Pic(dx, dy int) [][]uint8 {

	pic := make([][]uint8, dy)

	for y := range pic {
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			pic[y][x] = uint8((x + y) / 2)
		}
	}

	return pic
}

func main() {
	pic.Show(Pic)
}
