package main

import "fmt"

func main() {
	// 배열의 선언
	var arr [5]int // 크기가 5인 정수형 배열 선언

	// 배열의 초기화
	arr = [5]int{1, 2, 3, 4, 5} // 배열 요소 초기화

	// 배열 요소에 접근
	fmt.Println(arr[0]) // 첫 번째 요소에 접근하여 출력

	fmt.Println(len(arr)) // 배열의 길이 출력, 출력: 5

	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2) // 첫 번째 요소에 접근하여 출력

	var arr3 [5]int
	arr3[0] = 1
	arr3[3] = 4
	fmt.Println(arr3) // 첫 번째 요소에 접근하여 출력

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
