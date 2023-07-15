package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//標準入力から受け付ける入力回数
	n := 2
	sc := bufio.NewScanner(os.Stdin)
	//固定長nの整数値の配列を作成
	var x = make([]int, n)

	for i := 0; i < n; i++ {
		sc.Scan()
		x[i], _ = strconv.Atoi(sc.Text())
	}
	//入力値の積を出力
	fmt.Println(x[0] * x[1])
}
