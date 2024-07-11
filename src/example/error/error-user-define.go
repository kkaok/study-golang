package main

import (
	"fmt"
)

type MyError struct {
	Message string
	Code    int
}

func (e MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func doSomething(flag bool) error {
	if !flag {
		return MyError{
			Message: "something went wrong",
			Code:    123,
		}
	}
	return nil
}

func main() {
	err := doSomething(false)
	if err != nil {
		fmt.Println(err) // Error 123: something went wrong
	}
}
