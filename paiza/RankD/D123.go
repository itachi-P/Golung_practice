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
		s, _ := strconv.Atoi(sc.Text())
		if s < 10000 {
			s = s + 10000
		}
		fmt.Println(s)
	}
}
