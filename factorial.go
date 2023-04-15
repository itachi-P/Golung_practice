package main

import "fmt"

// 階乗を計算する再帰関数
func factorial(n int) int {
    if n == 0 { // nが0の場合、1を返す
        return 1
    }
    // nが0でない場合、n * (n-1)!を計算する
    return n * factorial(n-1)
}

func main() {
    var n int
    fmt.Print("階乗を計算する数を入力してください: ")
    fmt.Scan(&n) // 標準入力から数値を読み取る

    result := factorial(n) // 階乗を計算する

    fmt.Printf("%dの階乗は%dです\n", n, result)
}
