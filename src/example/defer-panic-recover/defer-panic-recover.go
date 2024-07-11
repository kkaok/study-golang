package main

import "fmt"

func riskyOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in riskyOperation:", r)
		}
	}()

	fmt.Println("Performing risky operation")
	panic("Something went wrong!")
	fmt.Println("This will not be printed")
}

func main() {
	fmt.Println("Start main")
	riskyOperation()
	fmt.Println("End main")
}
