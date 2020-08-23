package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "rancho", "the name of greeting object")
}

func main() {
	flag.Parse()
	fmt.Printf("hello, %s\n", name)
}
