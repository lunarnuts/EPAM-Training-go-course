package fibonacci

import (
	"fmt"
	"log"
)

func Printer(n int) {
	mp := recursive(n)
	for _, val := range mp {
		fmt.Printf("%d ", val)
	}
}

func recursive(n int) []int {
	if n < 0 {
		log.Fatal("number must be positive")
		return []int{}
	}
	if n == 0 {
		return []int{0}
	}
	mp := make([]int, n+1)
	for j := range mp {
		mp[j] = -1
	}
	mp[0], mp[1] = 0, 1
	recurUtil(n, mp)
	return mp
}

func recurUtil(n int, mp []int) int {
	if mp[n] != -1 {
		return mp[n]
	}
	result := recurUtil(n-1, mp) + recurUtil(n-2, mp)
	mp[n] = result
	return result
}
