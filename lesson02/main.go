package main

import (
	. "github.com/lunarnuts/go-course/tree/lesson02/fibonacci"
	"fmt"
)

func main() {
	fmt.Println("Fibonacci numbers is a form of sequence, such that each number is the sum of two preceeding ones.")
	Printer(40)
	fmt.Println()
	Recursive(40)
}

