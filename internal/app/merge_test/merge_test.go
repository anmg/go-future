package merge

import (
    "fmt"
    "testing"
)

func merge(li []int, left, mid, right int){
    i := left
    j := mid + 1
    tmp := []int{}
    for i <= mid && j <= right{
        if li[i] <= li[j]{
            tmp = append(tmp, li[i])
            i ++
        }else{
            tmp = append(tmp, li[j])
            j++
        }
    }

    if i <= mid {
        tmp = append(tmp, li[i:mid+1]...)
    }else{
        tmp = append(tmp, li[j:right+1]...)
    }
    for k := 0; k < len(tmp); k++{
        li[left+1] = tmp[k]
    }
}

func merge_sort(li []int, left, right int){
    if left < right {
       mid := (left + right)/2
       merge_sort(li, left, mid)
       merge_sort(li, mid+1, right)
       merge(li, left, mid, right)
    }
}

func TestMerge(t *testing.T){
    fmt.Println("hello")
}
