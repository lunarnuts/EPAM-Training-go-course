package main

import (
	. "github.com/lunarnuts/go-course/tree/lesson02/fibonacci"
	"fmt"
)

func main() {
	defer Recursive(20)
	defer fmt.Println()
	defer Printer(20)
	fmt.Println("Fibonacci numbers is a form of sequence, such that")
	fmt.Println("each number is the sum of two preceeding ones.\n")
}

