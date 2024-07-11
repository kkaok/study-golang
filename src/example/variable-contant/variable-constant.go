package main

import (
	"fmt"
)

func main() {
	// 변수 선언
	var x int
	x = 5
	fmt.Println("x:", x)

	// 변수 선언과 초기화
	var y int = 10
	fmt.Println("y:", y)

	// 타입 생략
	var z = 20
	fmt.Println("z:", z)

	// 짧은 선언
	a := 30
	fmt.Println("a:", a)

	// 여러 변수 선언
	var b, c, d int
	b, c, d = 1, 2, 3
	fmt.Println("b, c, d:", b, c, d)

	// 상수 선언
	const pi = 3.14
	fmt.Println("pi:", pi)

	// 여러 상수 선언
	const (
		width  = 100
		height = 200
	)
	fmt.Println("width, height:", width, height)

	// iota 사용
	const (
		e = iota // 0
		f        // 1
		g        // 2
	)
	fmt.Println("e, f, g:", e, f, g)
}
