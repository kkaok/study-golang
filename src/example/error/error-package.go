package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func findItem(id int) (string, error) {
	if id != 1 {
		return "", ErrNotFound
	}
	return "Item 1", nil
}

func main() {
	_, err := findItem(2)
	if err != nil {
		// Item not found
		if errors.Is(err, ErrNotFound) {
			fmt.Println("Item not found")
		} else {
			fmt.Println("An error occurred:", err)
		}
	}
}
