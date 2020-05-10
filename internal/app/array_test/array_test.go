package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var a [3]int
	a[0] = 1
	arr1 := [4]int{1, 2, 4, 5}
	arr2 := [...]int{1, 2, 4, 5}
	c := [2][2]int{{1, 2}, {4, 5}}
	t.Log(a, arr1, arr2, c)
	//b := [3] int {4,5,7}
	//c := [2][2] int {{1,2},{4,5}}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3_s := arr3[1:]
	t.Log(arr3_s)
	t.Log(arr3[1:])
}

func TestArrayTravel(t *testing.T){
	arr3 := [...] int {1,2,3,4,5}
	//t.Log(arr3[2:5])
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	for _, e := range arr3 {
		t.Log(e)
	}

	c := [2][2] int {{1,2},{4,5}}
	for i, e := range(c) {
		t.Log(i,e)
		for j, k := range(e) {
			t.Log(j, k)
		}
	}

	d := map[int] int{4:44}
	t.Log(d)
}
