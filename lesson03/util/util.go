package util

import (
	"fmt"
	"sort"
)

func FindAverage(arr []int) float64 {
	if len(arr) == 0 {
		return 0
	}
	sum := 0.0
	for _, elem := range arr {
		sum += float64(elem)
	}
	return sum / float64(len(arr))
}

func Max(sarr []string) string {
	if len(sarr) == 0 {
		return ""
	}
	maxstr := 0
	for ind, val := range sarr {
		if len(sarr[maxstr]) < len(val) {
			maxstr = ind
		}
	}
	return sarr[maxstr]
}

func Reverse(arr []int64) []int64 {
	l := len(arr)
	cp := make([]int64, l)
	for ind, val := range arr {
		cp[l-1-ind] = val
	}
	return cp
}

func PrintSorted(mp map[int]string) {
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
