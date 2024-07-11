package main

import (
	"fmt"
	"log"
	"os"
)

func riskyOperation() error {
	return fmt.Errorf("something went wrong")
}

func main() {
	err := riskyOperation()
	if err != nil {
		log.Printf("Operation failed: %v", err)
		os.Exit(1)
	}

	fmt.Println("Operation succeeded")
}
