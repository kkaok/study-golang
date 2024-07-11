package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Start")
	panic("Something went wrong!")
	fmt.Println("End") // 이 코드는 실행되지 않습니다.
}
