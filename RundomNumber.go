package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num int

	for i:= 1; i <= 5; i++ {
		fmt.Printf("おみくじ%d回目：", i)
		num = rand.Intn(6)
		switch num {
		case 0:
			fmt.Println("凶")
		case 1, 3:
			fmt.Println("吉")
		case 2, 4:
			fmt.Println("中吉")
		case 5:
			fmt.Println("大吉！")
		}
	}
}