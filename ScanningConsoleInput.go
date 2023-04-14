package main

import "fmt"

func main() {
	var input string
	fmt.Println("文字列を入力してください。")

	fmt.Scan(&input)

	fmt.Printf("%sと入力されました。", input)
}