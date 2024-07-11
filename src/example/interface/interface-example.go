package main

import "fmt"

// Shape 인터페이스 정의
type Shape interface {
	Area() float64
}

// Circle 구조체 정의
type Circle struct {
	Radius float64
}

// Circle 타입에 대한 Area 메서드 구현
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Rectangle 구조체 정의
type Rectangle struct {
	Width, Height float64
}

// Rectangle 타입에 대한 Area 메서드 구현
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 3}

	printArea(c) // 출력: Area: 78.5
	printArea(r) // 출력: Area: 12
}
