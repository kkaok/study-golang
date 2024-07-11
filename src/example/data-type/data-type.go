package main

import (
	"fmt"
)

// PersonWithName 구조체 정의
type PersonWithName struct {
	Name string
}

// Speak 메서드 정의
func (p PersonWithName) Speak() string {
	return "Hello, my name is " + p.Name
}

func main() {
	// 부울 타입
	var isGoAwesome bool = true
	fmt.Println("isGoAwesome:", isGoAwesome)

	// 정수 타입
	var age int = 30
	var height uint = 180
	var b byte = 255
	fmt.Println("age:", age, "height:", height, "byte:", b)

	// 부동 소수점 타입
	var pi float64 = 3.14
	fmt.Println("pi:", pi)

	// 복소수 타입
	var c complex128 = complex(1, 2)
	fmt.Println("complex number:", c)

	// 문자열 타입
	var greeting string = "Hello, World!"
	fmt.Println("greeting:", greeting)

	// 바이트와 루너 타입
	var letter byte = 'a'
	var unicodeChar rune = '한'
	fmt.Println("letter:", letter, "unicodeChar:", unicodeChar)

	// 배열 타입
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("numbers:", numbers)

	// 슬라이스 타입
	var scores []int = []int{10, 20, 30}
	scores = append(scores, 40)
	fmt.Println("scores:", scores)

	// 맵 타입
	var person map[string]int = map[string]int{"age": 30, "height": 180}
	fmt.Println("person:", person)

	// 구조체 타입
	type Person struct {
		Name string
		Age  int
	}
	var p Person = Person{Name: "John", Age: 30}
	fmt.Println("Person:", p)

	// 포인터 타입
	var x int = 10
	var pPointer *int = &x
	fmt.Println("pointer:", pPointer, "value:", *pPointer)

	// 인터페이스 타입
	type Speaker interface {
		Speak() string
	}

	var s Speaker = PersonWithName{Name: "John"}
	fmt.Println("Speaker:", s.Speak())

	// 함수 타입
	add := func(a int, b int) int {
		return a + b
	}
	var sum func(int, int) int = add
	fmt.Println("sum:", sum(1, 2))
}
