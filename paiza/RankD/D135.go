package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//標準入力から受け付ける入力回数
	n := 1
	sc := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		sc.Scan()
		n, _ := strconv.Atoi(sc.Text())
		//x角形の内角の和が与えられる
		x := n/180 + 2
		fmt.Println(x)
	}
}
