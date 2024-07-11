package main

import "fmt"

func main() {
	// 맵 초기화
	ages := map[string]int{
		"Alice":   31,
		"Bob":     34,
		"Charlie": 28,
	}

	// 요소 추가
	ages["David"] = 23

	// 요소 조회
	fmt.Println("Alice's age is", ages["Alice"])

	// 요소 삭제
	delete(ages, "Bob")

	// 키 체크
	if age, ok := ages["Charlie"]; ok {
		fmt.Println("Charlie's age is", age)
	} else {
		fmt.Println("Charlie's age is not present")
	}

	// 맵 순회
	for name, age := range ages {
		fmt.Printf("%s's age is %d\n", name, age)
	}
}
