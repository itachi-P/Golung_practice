package main

import "fmt"

func main() {
	name := "Blacksmith"

	fmt.Println(name)
	//変数のポインタを別の関数に渡して値を変更
	changeStr(&name)

	fmt.Println(name)

}

func changeStr(strPtr *string) {
	*strPtr = "Carpenter"

}
