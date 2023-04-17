package main

import "fmt"

// fibonacciはn番目のフィボナッチ数を返します。
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Print("Enter the number of terms: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print(fibonacci(i), " ")
	}
}
