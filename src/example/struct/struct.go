package main

import "fmt"

// 구조체 정의
type Person struct {
	Name string
	Age  int
}

// 구조체 메서드 정의
func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

func (p *Person) UpdateName(newName string) {
	p.Name = newName
}

type Employee struct {
	Person
	Position string
}

func main() {
	// 구조체 초기화
	p := Person{
		Name: "Alice",
		Age:  30,
	}

	// 필드 접근
	fmt.Println("Name:", p.Name) // 출력: Name: Alice
	fmt.Println("Age:", p.Age)   // 출력: Age: 30

	// 메서드 호출
	fmt.Println(p.Greet()) // 출력: Hello, my name is Alice

	p2 := Person{"Alice", 30}
	p2.UpdateName("Bob")
	fmt.Println(p2.Name) // 출력: Bob

	e := Employee{
		Person:   Person{Name: "Alice", Age: 30},
		Position: "Developer",
	}
	fmt.Println(e.Name)     // 출력: Alice
	fmt.Println(e.Age)      // 출력: 30
	fmt.Println(e.Position) // 출력: Developer
}
