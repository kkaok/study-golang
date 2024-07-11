package main

import (
	"errors"
	"fmt"
)

func doSomething() error {
	return fmt.Errorf("doSomething failed: %w", errors.New("underlying error"))
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println(err)

		var underlyingError error
		// doSomething failed: underlying error
		// Underlying error: doSomething failed: underlying error
		if errors.As(err, &underlyingError) {
			fmt.Println("Underlying error:", underlyingError)
		}
	}
}
