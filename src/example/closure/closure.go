package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	c := counter()
	fmt.Println(c()) // 1 출력
	fmt.Println(c()) // 2 출력
	fmt.Println(c()) // 3 출력
}
