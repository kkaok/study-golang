package main

import "fmt"

// Circle 구조체 정의
type Circle struct {
	Radius float64
}

// Area 메서드 정의 (값 리시버 사용)
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Scale 메서드 정의 (포인터 리시버 사용)
func (c *Circle) Scale(factor float64) {
	c.Radius *= factor
}

func main() {
	c := Circle{Radius: 5}
	fmt.Println("Original Radius:", c.Radius) // 출력: Original Radius: 5
	fmt.Println("Area:", c.Area())            // 출력: Area: 78.5

	c.Scale(2)
	fmt.Println("Scaled Radius:", c.Radius) // 출력: Scaled Radius: 10
	fmt.Println("Area:", c.Area())          // 출력: Area: 314
}
