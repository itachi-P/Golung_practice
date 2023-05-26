package gosample

import "fmt"

var Message string = "Hello, %s!\n"

func Hello(name string) {
	fmt.Printf(Message, name)
}
