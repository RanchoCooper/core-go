package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	example1()
	example2()
}

func example1() {
	intChannels := [3]chan int {
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	index := rand.Intn(3)
	fmt.Printf("the index: %d\n", index)
	intChannels[index] <- index

	select {
	case <-intChannels[0]:
		fmt.Println("the first candidate case is selected")
	case <-intChannels[1]:
		fmt.Println("the second candidate case is selected")
	case <-intChannels[2]:
		fmt.Println("the third candidate case is selected")
	default:
		fmt.Println("no candidate case is selected")
	}
}

func example2() {
	intChan := make(chan int, 1)
	time.AfterFunc(time.Second, func() {
		close(intChan)
	})

	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Println("the candidate case is closed.")
			break
		}
		fmt.Println("the candidate case is selected")
	}
}