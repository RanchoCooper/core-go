package main

import (
	"fmt"
)

var container = []string{"zero", "one", "two"}

func main() {
	value, ok := interface{}(container).([]string)
	if ok {
		fmt.Printf("The element is %q.\n", value)
	}
	// container := map[int]string{0: "zero", 1: "one", 2: "two"}
	// // value, ok := interface{}(container).(map[int]string)
	// // if ok {
	// // 	fmt.Printf("The element is %q.\n", value)
	// }
}
