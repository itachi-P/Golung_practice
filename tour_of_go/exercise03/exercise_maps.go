package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	fmt.Printf("%q\n", words)
	m := make(map[string]int)

	for _, v := range words {
		count := 0
		if v, ok := m[v]; ok {
			count = v
		}
		m[v] = count + 1

		fmt.Println(v)
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
