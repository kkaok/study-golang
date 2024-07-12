package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // 채널 닫기
	}()

	for i := range ch {
		fmt.Println(i)
	}
}
