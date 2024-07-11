package main

import "fmt"

// 두 정수를 더하는 함수
func add(x int, y int) int {
	return x + y
}

// 문자열을 출력하는 함수
func greet(name string) {
	fmt.Println("Hello, " + name + "!")
}

// 두 개의 정수와 두 개의 문자열을 반환하는 함수
func complexFunction(a int, b int) (int, int, string, string) {
	sum := a + b
	difference := a - b
	return sum, difference, "Sum", "Difference"
}

func main() {
	// 함수 호출
	result := add(3, 5)
	fmt.Println("3 + 5 =", result)

	greet("Alice")

	sum, diff, sumLabel, diffLabel := complexFunction(10, 5)
	fmt.Println(sumLabel, ":", sum)
	fmt.Println(diffLabel, ":", diff)
}
