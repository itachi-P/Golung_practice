package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//標準入力から任意の数値を読み込み、指定のルールに従った簡易二次元バーコードとして表現する。
//入力される値Nは整数値1個で、入力の末尾に改行が1個入る。
//桁数は3の倍数かつ100<=N<=999,999,999の範囲
//値は文字列で標準入力から渡される。

// 標準入力から文字列を1行取得する関数
func StrStdin() (stringInput string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringInput = scanner.Text()
	stringInput = strings.TrimSpace(stringInput)
	return
}

// 与えられた1桁の数値を3×3マスの二次元バーコードに変換し、二次元配列に格納する
func NumToBarcode(num int) [3][3]string {
	// 3×3マスで表現される二次元バーコードを格納する二次元配列
	var barcode [3][3]string
	for index := 1; index <= 9; {
		for row := 0; row < 3; row++ {
			for col := 0; col < 3; col++ {
				if num < index {
					barcode[row][col] = "."
				} else {
					barcode[row][col] = "#"
				}
				index++
			}
		}
	}
	return barcode
}

// 渡された3文字を3桁の数値に変換し、3×3マス×3ブロックの二次元バーコードとして出力する
func ConvBarcodeBlock(input string) {
	// 文字列を数値に変換
	num, _ := strconv.Atoi(input)
	// 3桁の数値を3×3マスの二次元バーコード×3ブロックに変換し、標準出力に表示
	num1 := NumToBarcode(num / 100)
	num2 := NumToBarcode(num / 10 % 10)
	num3 := NumToBarcode(num % 10)
	PrintBarcodeBlock(num1, num2, num3)
}

// 3×3マス×3ブロックの二次元バーコードを標準出力に表示する
func PrintBarcodeBlock(num1, num2, num3 [3][3]string) {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			fmt.Print(num1[row][col])
		}
		for col := 0; col < 3; col++ {
			fmt.Print(num2[row][col])
		}
		for col := 0; col < 3; col++ {
			fmt.Print(num3[row][col])
		}
		fmt.Println()
	}
}

func main() {
	// 標準入力から文字列を1行入力
	input := StrStdin()

	// 前から順に3文字づつに分割し、バーコードに変換する
	for i := 0; i < len(input); i += 3 {
		ConvBarcodeBlock(input[i : i+3])
	}
}
