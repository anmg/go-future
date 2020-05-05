package sort_test

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

//冒泡排序
func bubble_sort(li []int) {
	for i := 0; i < len(li)-1; i++ {
		exchange := false
		for j := 0; j < len(li)-i-1; j++ {
			if li[j] > li[j+1] {
				li[j], li[j+1] = li[j+1], li[j]
				exchange = true
			}
		}
		if !exchange {
			return
		}
	}
}

//选择排序
func select_sort(li []int) {
	for i := 0; i < len(li)-1; i++ {
		pos := i
		for j := i + 1; j < len(li); j++ {
			if li[pos] > li[j] {
				pos = j
			}
		}
		li[i], li[pos] = li[pos], li[i]
	}
}

//插入排序
func insert_sort(li []int) {
	for i := 1; i < len(li); i++ {
		tmp := li[i]
		j := i - 1
		for j >= 0 && tmp < li[j] {
			li[j+1] = li[j]
			j --
		}
		li[j+1] = tmp
	}
}

//希尔排序
func shell_sort(li []int) {
	for gap := len(li) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(li); i++ {
			tmp := li[i]
			j := i - gap
			for j >= 0 && tmp < li[j] {
				li[j+gap] = li[j]
				j -= gap
			}
			li[j+gap] = tmp
		}
	}
}

//快速排序
func quick_sort(li []int, left, right int) {
	if left >= right {
		return
	}
	i := left
	j := right
	//rand.Seed(time.Now().Unix())
	//r := rand.Intn(right-left) + left
	//li[i], li[r] = li[r], li[i]

	tmp := li[i]
	for i < j {
		for i < j && li[j] >= tmp {
			j--
		}
		li[i] = li[j]
		for i < j && li[i] <= tmp {
			i++
		}
		li[j] = li[i]
	}
	li[i] = tmp
	quick_sort(li, left, i-1)
	quick_sort(li, i+1, right)
}

//堆排序
func sift(li []int, low, high int) {
	i := low
	j := 2*i + 1
	tmp:=li[i]
	for j <= high {
		if j < high && li[j] < li[j+1] {
			j++
		}
		if tmp < li[j] {
			li[i]  = li[j]
			i = j
			j = 2*i + 1
		} else {
			break
		}
	}
	li[i] = tmp
}

func heap_sort(li []int) {
	for i := len(li)/2 - 1; i >= 0; i-- {
		sift(li, i, len(li)-1)
	}
	for j := len(li) - 1; j > 0; j-- {
		li[0], li[j] = li[j], li[0]
		sift(li, 0, j-1)
	}
}

//归并排序
func merge(li []int, left, mid, right int) {
	i := left
	j := mid + 1
	tmp := []int{}
	for i <= mid && j <= right {
		if li[i] <= li[j] {
			tmp = append(tmp, li[i])
			i ++
		} else {
			tmp = append(tmp, li[j])
			j ++
		}
	}
	if i <= mid {
		tmp = append(tmp, li[i:mid+1]...)
	} else {
		tmp = append(tmp, li[j:right+1]...)
	}
	for k := 0; k < len(tmp); k++ {
		li[left+k] = tmp[k]
	}
}

func merge_sort(li []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		merge_sort(li, left, mid)
		merge_sort(li, mid+1, right)
		merge(li, left, mid, right)
	}
}

//计数排序
func count_sort(li []int) {
   max_num := li[0]
   for i := 1; i < len(li); i++ {
       if max_num < li[i] {
           max_num = li[i]
       }
   }
   arr := make([]int, max_num+1)
   for j := 0; j < len(li); j++ {
       arr[li[j]]++
   }
   k := 0
   for m, n := range arr {
       for p := 0; p < n; p++ {
           li[k] = m
           k++
       }
   }
}

//桶排序
func bin_sort(li []int, bin_num int) {
	min_num, max_num := li[0], li[0]
	for i := 0; i < len(li); i++ {
		if min_num > li[i] {
			min_num = li[i]
		}
		if max_num < li[i] {
			max_num = li[i]
		}
	}
	bin := make([][]int, bin_num)
	for j := 0; j < len(li); j++ {
		n := (li[j] - min_num) / ((max_num - min_num + 1) / bin_num)
		bin[n] = append(bin[n], li[j])
		k := len(bin[n]) - 2
		for k >= 0 && li[j] < bin[n][k] {
			bin[n][k+1] = bin[n][k]
			k--
		}
		bin[n][k+1] = li[j]
	}
	o := 0
	for p, q := range bin {
		for t := 0; t < len(q); t++ {
			li[o] = bin[p][t]
			o++
		}
	}
}

//基数排序
func radix_sort(li []int) {
	max_num := li[0]
	for i := 0; i < len(li); i++ {
		if max_num < li[i] {
			max_num = li[i]
		}
	}
	for j := 0; j < len(strconv.Itoa(max_num)); j++ {
		bin := make([][]int, 10)
		for k := 0; k < len(li); k++ {
			n := li[k] / int(math.Pow(10, float64(j))) % 10
			bin[n] = append(bin[n], li[k])
		}
		m := 0
		for p := 0; p < len(bin); p++ {
			for q := 0; q < len(bin[p]); q++ {
				li[m] = bin[p][q]
				m++
			}
		}
	}
}

func TestSort(t *testing.T)  {
	//nums := []int{-1, 0, 1, 2, -1, -4}
	nums := []int{2,4,1,9,5,7,3}
	fmt.Println("input:", nums)
	quick_sort(nums, 0, 6)
	fmt.Println("quick_sort", nums)

	bubble_sort(nums)
	fmt.Println("bubble_sort", nums)
}

