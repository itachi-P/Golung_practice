package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	x, y, tmp := 0, 1, 0
	return func(i int) int {
		if i < 2 {
			return i
		}
		tmp = x + y
		x = y
		y = tmp
		return tmp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
