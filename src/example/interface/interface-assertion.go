package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s, ok := i.(string)
	if ok {
		fmt.Println(s) // 출력: hello
	} else {
		fmt.Println("i is not a string")
	}
}
