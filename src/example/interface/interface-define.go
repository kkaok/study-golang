package main

import "fmt"

// Speaker 인터페이스 정의
type Speaker interface {
	Speak() string
}

// Person 구조체 정의
type Person struct {
	Name string
}

// Person 타입에 대한 Speak 메서드 구현
func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

// Dog 구조체 정의
type Dog struct {
	Name string
}

// Dog 타입에 대한 Speak 메서드 구현
func (d Dog) Speak() string {
	return "Woof! My name is " + d.Name
}

func main() {
	var s Speaker

	p := Person{Name: "Alice"}
	d := Dog{Name: "Buddy"}

	s = p
	fmt.Println(s.Speak()) // 출력: Hello, my name is Alice

	s = d
	fmt.Println(s.Speak()) // 출력: Woof! My name is Buddy
}
