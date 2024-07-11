package main

import (
	"fmt"
)

func main() {
	x := 10

	// if 문
	if x > 5 {
		fmt.Println("x는 5보다 큽니다.")
	}

	// if-else 문
	if x < 5 {
		fmt.Println("x는 5보다 작습니다.")
	} else {
		fmt.Println("x는 5보다 크거나 같습니다.")
	}

	// if-else if-else 문
	if x < 5 {
		fmt.Println("x는 5보다 작습니다.")
	} else if x == 10 {
		fmt.Println("x는 10입니다.")
	} else {
		fmt.Println("x는 5보다 크고 10은 아닙니다.")
	}

	// 초기화 문이 있는 if 문
	if y := 20; y > 15 {
		fmt.Println("y는 15보다 큽니다.")
	}

	// switch 문
	switch x {
	case 1:
		fmt.Println("x는 1입니다.")
	case 2:
		fmt.Println("x는 2입니다.")
	case 10:
		fmt.Println("x는 10입니다.")
	default:
		fmt.Println("x는 1, 2, 10 중 하나가 아닙니다.")
	}

	// 여러 조건을 가진 switch 문
	switch x {
	case 1, 2, 3:
		fmt.Println("x는 1, 2, 3 중 하나입니다.")
	case 10, 20, 30:
		fmt.Println("x는 10, 20, 30 중 하나입니다.")
	default:
		fmt.Println("x는 1, 2, 3, 10, 20, 30 중 하나가 아닙니다.")
	}

	// 조건이 있는 switch 문
	switch {
	case x < 5:
		fmt.Println("x는 5보다 작습니다.")
	case x >= 5 && x < 15:
		fmt.Println("x는 5 이상 15 미만입니다.")
	default:
		fmt.Println("x는 15 이상입니다.")
	}
}
