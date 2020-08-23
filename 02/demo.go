package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
	flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	flag.CommandLine.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.StringVar(&name, "name", "rancho", "the name of greeting object")
}

func main() {
	// flag.Usage = func() {
	// 	_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }

	// flag.Parse()
	cmdLine.Parse(os.Args[1:])

	fmt.Printf("hello, %s\n", name)
}
