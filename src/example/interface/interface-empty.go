package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", i, i)
}

func main() {
	describe(42)      // Value: 42, Type: int
	describe("hello") // Value: hello, Type: string
	describe(true)    // Value: true, Type: bool
}
