package main

import (
	"fmt"
)

func main() {
	var badMap = map[interface{}]int {
		"1": 1,
		[]int{2}: 2,
		3: 3,
	}
	// panic: runtime error: hash of unhashable type []int
	fmt.Printf("%v", badMap)
}
