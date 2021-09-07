package main

import (
	"fmt"

	"github.com/lunarnuts/go-course/tree/lesson02/fibonacci"
)

func main() {
	defer fmt.Println()
	defer fibonacci.Printer(20)
	fmt.Println("Fibonacci numbers is a form of sequence, such that")
	fmt.Println("each number is the sum of two preceeding ones.")
}
