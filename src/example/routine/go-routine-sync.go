package main

import (
	"fmt"
)

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	fmt.Println(a)
	fmt.Println(total)
	c <- total // total 값을 채널에 전송
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make(chan int)

	//fmt.Println(a[:len(a)/2])
	//fmt.Println(a[len(a)/2:])

	go sum(a[len(a)/2:], c)
	go sum(a[:len(a)/2], c)
	x, y := <-c, <-c // 채널로부터 값을 수신
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x, y, x+y)
}
