package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("enter function main")
	defer func() {
		fmt.Println("enter defer function")
		if p := recover(); p != nil {
			fmt.Printf("panic: %s", p)
		}
		fmt.Println("exit defer function")
	}()

	panic(errors.New("something wrong\n"))

	// unreachable code
	fmt.Println("exit function main")
}
