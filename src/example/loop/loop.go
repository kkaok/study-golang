package main

import "fmt"

func main() {
	// 기본 for 문
	for i := 0; i < 5; i++ {
		fmt.Println("기본 for 문:", i)
	}

	// 조건만 있는 for 문
	j := 0
	for j < 5 {
		fmt.Println("조건만 있는 for 문:", j)
		j++
	}

	// 무한 루프
	k := 0
	for {
		if k >= 5 {
			break
		}
		fmt.Println("무한 루프:", k)
		k++
	}

	// 슬라이스와 배열을 순회하는 for 문
	nums := []int{2, 4, 6, 8, 10}
	for index, value := range nums {
		fmt.Printf("슬라이스: index=%d, value=%d\n", index, value)
	}

	// 문자열을 순회하는 for 문
	str := "hello"
	for index, char := range str {
		fmt.Printf("문자열: index=%d, char=%c\n", index, char)
	}

	// 맵을 순회하는 for 문
	fruits := map[string]string{"a": "apple", "b": "banana", "c": "cherry"}
	for key, value := range fruits {
		fmt.Printf("맵: key=%s, value=%s\n", key, value)
	}

	// break 예제
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("break 예제:", i)
	}

	// continue 예제
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("continue 예제:", i)
	}
}
