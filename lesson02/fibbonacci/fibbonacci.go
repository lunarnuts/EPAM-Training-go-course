package fibbonacci

import "fmt"

func Printer(n int) {
	one, two := 0, 1
	fmt.Print(one, " ", two)
	for n > 2 {
		fib := one + two
		one, two = two, fib
		n--
		fmt.Printf(" %d", fib)
	}
}

func PrinterRecursive(n int) int {
	if n == 0 {
		fmt.Print(0)
		return 0
	}
	if n == 1 {
		fmt.Print(" ", 1)
		return 1
	}
	result := PrinterRecursive(n-1) + PrinterRecursive(n-2)
	fmt.Print(" ", result)
	return result
}
