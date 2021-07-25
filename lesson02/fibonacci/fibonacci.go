package fibonacci

import (
	"fmt"
	"strings"
)

func Printer(n int) {
	one := 0
	two := 1
	fmt.Print(one, " ", two, " ")
	for n > 1 {
		fib := one + two
		one, two = two, fib
		n--
		fmt.Print(fib, " ")
	}
	fmt.Println()
}

func Recursive(n int) {
	if n < 0 {
		fmt.Println("Number must be positive")
		return
	}
	mp := make([]int, n+1)
	for j := range mp {
		mp[j] = -1
	}
	mp[0], mp[1] = 0, 1
	recurUtil(n, mp)
	var b strings.Builder
	for _, i := range mp {
		fmt.Fprintf(&b, "%d ", i)
	}
	fmt.Println(b.String())
}
func recurUtil(n int, mp []int) int {
	if mp[n] != -1 {
		return mp[n]
	}
	result := recurUtil(n-1, mp) + recurUtil(n-2, mp)
	mp[n] = result
	return result
}
