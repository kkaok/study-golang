package main

import "fmt"

func main() {
	// 슬라이스 선언
	var slice1 []int
	// 슬라이스 초기화
	slice1 = []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)
	// 선언과 초기화를 동시에 하기
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice2)

	// 길이
	fmt.Println("Length of slice:", len(slice2)) // 출력: 5
	// 용량
	fmt.Println("Capacity of slice:", cap(slice2)) // 출력: 5 (기본적으로 길이와 동일)

	// 부분 슬라이스 생성
	subSlice := slice2[1:4]             // 인덱스 1부터 3까지의 부분 슬라이스
	fmt.Println("subSlice :", subSlice) // 출력: 5

	// 요소 추가
	slice2 = append(slice2, 6, 7) // 슬라이스에 4와 5 추가
	fmt.Println("slice add : ", slice2)
}
