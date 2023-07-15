package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//標準入力で与えられる文字列の周辺を記号で囲む

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("文字列を入力してください: ")
	scanner.Scan()
	str := scanner.Text()
	length := len(str)

	//上部、文字列の段、下部の順に出力
	fmt.Println("+" + strings.Repeat("-", length) + "+")
	fmt.Printf("|%s|\n", str)
	fmt.Println("+" + strings.Repeat("-", length) + "+")
}
