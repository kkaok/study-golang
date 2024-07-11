package main

import (
	"fmt"
	"mymath" // 패키지 임포트
)

func main() {
	result1 := mymath.Add(10, 5)
	fmt.Println("10 + 5 =", result1) // 출력: 10 + 5 = 15

	result2 := mymath.Subtract(10, 5)
	fmt.Println("10 - 5 =", result2) // 출력: 10 - 5 = 5
}
