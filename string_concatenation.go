package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//文字列の連結
//最初に入力回数が与えられ、2行目以降で入力された文字列から空白を削除する

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		str := scanner.Text()
		str = strings.Replace(str, " ", "", -1) //-1は回数制限なく全て置換するという意味
		//strings.TrimSpace(str)//これだと前後の空白のみ削除される
		fmt.Println(str)
	}
}
