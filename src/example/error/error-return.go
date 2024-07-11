package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(4, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result) // result : 2
	}

	result, err = divide(4, 0)
	if err != nil {
		fmt.Println("Error:", err) // Error: division by zero
	} else {
		fmt.Println("Result:", result)
	}
}
