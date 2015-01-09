package main

import (
		"fmt"
		"github.com/jaredharley/learn-go/stringutil"
)

func main() {
	fmt.Print("Hello, world!\n")
	fmt.Printf(stringutil.Reverse("Hello, world!"))
	fmt.Print("\n")
}