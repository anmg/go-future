package sort_test

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

func TestSortReverse(t *testing.T)  {
	intList := [] int {2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	float8List := [] float64 {4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	stringList := [] string {"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}

	sort.Reverse(sort.IntSlice(intList))

	fmt.Println("--0--",intList)
	sort.Ints(intList)
	fmt.Println("--1--",intList)
	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	fmt.Println("--2--",intList)

	sort.Sort(sort.Reverse(sort.Float64Slice(float8List)))
	sort.Sort(sort.Reverse(sort.StringSlice(stringList)))

	//fmt.Printf("%v\n%v\n%v\n", intList, float8List, stringList)
}


func TestSort(t *testing.T)  {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println("--1--",nums)
	sort.Ints(nums)
	fmt.Println("--2--",nums)
	sort.IntsAreSorted(nums)
	fmt.Println("--3--",nums)


	aa := strconv.FormatInt(4343, 10)
	fmt.Println("--3--",aa)

}