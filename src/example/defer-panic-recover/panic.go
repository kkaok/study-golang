package main

import "fmt"

func main() {
	fmt.Println("Start")
	panic("Something went wrong!")
	fmt.Println("End") // 이 코드는 실행되지 않습니다.
}
