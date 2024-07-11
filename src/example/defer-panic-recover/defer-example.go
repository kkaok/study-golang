package main

import (
	"fmt"
	"os"
)

func readFile(fileName string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", err))
	}
	defer file.Close()

	// 파일 처리 로직 (예: 읽기)
	fmt.Println("File opened successfully")
}

func main() {
	readFile("example.txt") // Recovered from panic: Failed to open file: open example.txt: no such file or directory
	fmt.Println("Continuing execution")
}
