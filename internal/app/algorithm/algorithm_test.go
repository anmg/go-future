package sort_test

import (
	"fmt"
	"sort"
	"testing"
)

func twoSum1(nums[]int, target int) []int {
	ret := []int{}
	for i, num := range nums{
		temp := target - num
		//fmt.Println(temp)
		for j := i+1; j < len(nums); j++{
			//fmt.Println(nums[j],temp)
			if nums[j] == temp {
				ret = append(ret, i)
				ret = append(ret, j)
				break
			}
		}

		if len(ret) > 0 {
			break
		}
	}

	return ret
}

func twoSum(nums[]int, target int) [][]int {
	ret := [][]int{}
	ret_map := map[int]bool{}
	num_map := map[int]int{}

	for i, num := range nums{
		temp := target - num
		if _,ok := num_map[temp];ok {
			if ret_map[temp] != true && ret_map[num] != true {
				num_pairs := []int{}
				num_pairs = append(num_pairs, temp)
				num_pairs = append(num_pairs, num)
				ret = append(ret, num_pairs)
				ret_map[temp] = true
				ret_map[num] = true
			}
		}else{
			num_map[num] = i
		}

		//if len(ret) > 0 {
		//	break
		//}
	}

	return ret
}

func threeSum1(nums []int) [][]int {
	ret := [][]int{}
	for i, num := range nums{
		target := 0 - num
		num_map := map[int] int {}
		ret2 := []int{}
		for j := i + 1; j < len(nums); j++{
			temp := target - nums[j]
			if _,ok := num_map[temp];ok {
				ret2 = append(ret2, num)
				ret2 = append(ret2, temp)
				ret2 = append(ret2, nums[j])
			}else{
				num_map[nums[j]] = temp
			}
			if len(ret2) > 0 {
				ret = append(ret, ret2)
				ret2 = []int{}
				num_map = map[int] int {}
			}
		}

	}

	return ret
}

func threeSum(nums []int) [][]int {
	ret := [][]int{}
	fmt.Println("--1--",nums)
	sort.Ints(nums)
	fmt.Println("--2--",nums)
	sort.IntsAreSorted(nums)
	fmt.Println("--3--",nums)
	for i, num := range nums{
		target := 0 - num
		num_map := map[int] int {}
		ret2 := []int{}
		for j := i + 1; j < len(nums); j++{
			temp := target - nums[j]
			if _,ok := num_map[temp];ok {
				ret2 = append(ret2, num)
				ret2 = append(ret2, temp)
				ret2 = append(ret2, nums[j])
			}else{
				num_map[nums[j]] = temp
			}
			if len(ret2) > 0 {
				ret = append(ret, ret2)
				ret2 = []int{}
				num_map = map[int] int {}
			}
		}

	}

	return ret
}

//func TestThreeSum(t *testing.T)  {
//	nums := []int{-1, 0, 1, 2, -1, -4}
//	fmt.Println(nums)
//	ret := threeSum(nums)
//	fmt.Println(ret)
//}

func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 4, 6, 7, 1, 3, 5, 5, 1}
	fmt.Println(nums)
	ret := twoSum(nums, 6)
	fmt.Println(ret)
}