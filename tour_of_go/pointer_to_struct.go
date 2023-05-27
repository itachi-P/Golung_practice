package main

import "fmt"

//Pointers to structs

type Person struct {
	Name   string
	Age    int
	Height float32
}

func main() {
	p := Person{"Manji", 14, 165.7}
	pp := &p         //構造体もポインタで参照できる
	(*pp).Age = 2e3  //構造体の要素(Field)にアクセスする
	pp.Height = 1e-4 //上記を簡略化して書くこともできる
	fmt.Println(p)
}
