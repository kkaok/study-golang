package main

import (
	"fmt"
)

func main() {
	// 산술 연산자
	a := 10
	b := 5
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)

	// 증가 및 감소 연산자
	a++
	fmt.Println("a++ =", a)
	b--
	fmt.Println("b-- =", b)

	// 비교 연산자
	fmt.Println("a == b =", a == b)
	fmt.Println("a != b =", a != b)
	fmt.Println("a > b =", a > b)
	fmt.Println("a < b =", a < b)
	fmt.Println("a >= b =", a >= b)
	fmt.Println("a <= b =", a <= b)

	// 논리 연산자
	fmt.Println("true && false =", true && false)
	fmt.Println("true || false =", true || false)
	fmt.Println("!true =", !true)

	// 비트 연산자
	c := 6                        // 0110 in binary
	d := 11                       // 1011 in binary
	fmt.Println("c & d =", c&d)   // 0010 in binary, which is 2
	fmt.Println("c | d =", c|d)   // 1111 in binary, which is 15
	fmt.Println("c ^ d =", c^d)   // 1101 in binary, which is 13
	fmt.Println("c &^ d =", c&^d) // 0100 in binary, which is 4

	// 시프트 연산자
	e := 8                        // 1000 in binary
	fmt.Println("e << 1 =", e<<1) // 10000 in binary, which is 16
	fmt.Println("e >> 1 =", e>>1) // 0100 in binary, which is 4

	// 할당 연산자
	f := 5
	f += 3
	fmt.Println("f += 3 =", f)
	f -= 2
	fmt.Println("f -= 2 =", f)
	f *= 2
	fmt.Println("f *= 2 =", f)
	f /= 3
	fmt.Println("f /= 3 =", f)
	f %= 3
	fmt.Println("f %= 3 =", f)

	// 주소 및 포인터 연산자
	g := 42
	var p *int = &g
	fmt.Println("g의 주소:", p)
	fmt.Println("p가 가리키는 값:", *p)
}
