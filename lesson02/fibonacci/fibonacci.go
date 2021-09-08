package fibonacci

import (
	"fmt"
)

func Printer(n int) {
	mp, err := recursive(n)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	for _, val := range mp {
		fmt.Printf("%d ", val)
	}
}

func recursive(n int) ([]int, error) {
	if n < 0 {
		return []int{}, fmt.Errorf("number must be non-negative")
	}
	if n == 0 {
		return []int{0}, nil
	}
	mp := make([]int, n+1)
	for j := range mp {
		mp[j] = -1
	}
	mp[0], mp[1] = 0, 1
	recurUtil(n, mp)
	return mp, nil
}

func recurUtil(n int, mp []int) int {
	if mp[n] != -1 {
		return mp[n]
	}
	result := recurUtil(n-1, mp) + recurUtil(n-2, mp)
	mp[n] = result
	return result
}
