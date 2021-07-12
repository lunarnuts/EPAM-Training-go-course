package main

import (
	"fmt"
	"github.com/lunarnuts/go-course/tree/lesson03/util"
)

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
	fmt.Println(util.FindAverage(arr))
	fmt.Println(util.Max(sarr))
	fmt.Println(util.Reverse(ar))
	util.PrintSorted(mp)
}
