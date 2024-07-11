package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	// 복소수 선언
	var c1 complex64 = complex(5, 7)       // 5 + 7i
	var c2 complex128 = complex(3.2, -4.5) // 3.2 - 4.5i

	// 복소수의 실수 부분과 허수 부분 추출
	fmt.Println("c1:", c1)
	fmt.Println("실수 부분:", real(c1), "허수 부분:", imag(c1))

	fmt.Println("c2:", c2)
	fmt.Println("실수 부분:", real(c2), "허수 부분:", imag(c2))

	// 복소수 연산
	sum := c1 + complex64(c2)     // 복소수 덧셈
	product := c1 * complex64(c2) // 복소수 곱셈

	fmt.Println("합:", sum)
	fmt.Println("곱:", product)

	// 복소수 함수
	abs := cmplx.Abs(c2)     // 복소수의 절대값
	phase := cmplx.Phase(c2) // 복소수의 위상

	fmt.Println("c2의 절대값:", abs)
	fmt.Println("c2의 위상:", phase)
}
