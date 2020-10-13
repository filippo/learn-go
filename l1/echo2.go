package main

import (
	"fmt"
	"os"
)

func main() {
	// all the equivalent variable declaration
	// var s string
	// var s = ""
	// var s string = ""
	// s:= ""
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
