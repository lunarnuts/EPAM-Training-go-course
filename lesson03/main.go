package main

import (
	"fmt"
	"sort"
)

func mean(arr []int) float64 {
	if len(arr) == 0 {
		fmt.Println("empty array")
		return 1
	}
	sum := 0.0
	for _, elem := range arr {
		sum += float64(elem)
	}
	return sum / float64(len(arr))
}

func max(sarr []string) string {
	if len(sarr) == 0 {
		return "empty array"
	}
	maxstr := 0
	for ind, val := range sarr {
		if len(sarr[maxstr]) < len(val) {
			maxstr = ind
		}
	}
	return sarr[maxstr]
}

func reverse(arr []int64) []int64 {
	l := len(arr)
	cp := make([]int64, l)
	for ind, val := range arr {
		cp[l-1-ind] = val
	}
	return cp
}

func printSorted(mp map[int]string) {
	keys := []int{}
	values := []string{}
	for key, _ := range mp {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, key := range keys {
		values = append(values, mp[key])
	}
	fmt.Println(values)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	ar := []int64{9, 8, 7, 6, 5}
	sarr := []string{"hohohoh", "hello", "world!", "test", "longest"}
	mp := map[int]string{
		33: "hey",
		1:  "hello",
		10: "world",
		9:  "hohohoh",
		5:  "test",
	}
	fmt.Println(mean(arr))
	fmt.Println(max(sarr))
	fmt.Println(reverse(ar))
	printSorted(mp)
}
