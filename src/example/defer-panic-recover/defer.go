package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close() // 함수가 끝나기 직전에 파일을 닫습니다.

	// 파일 작업 수행
	fmt.Println("File opened successfully")
}
