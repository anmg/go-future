package slice_test

import (
	"testing"
)

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0,1)

	t.Log(len(s0), cap(s0))

	s1 := [] int {1,2,3,5,6}

	t.Log(s1, len(s1), cap(s1))

	s2 := make([]int, 3, 5)

	t.Log(s2, len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])

	s2 = append(s2, 9)
	s2 = append(s2, 10)
	s2 = append(s2, 10)
	//s2 = append(s2, 10)
	//s2 = append(s2, 10)

	t.Log(s2, len(s2), cap(s2))
	//t.Log(s2[0], s2[1], s2[2], s2[3], s2[4])
}

func TestSliceGroving(t *testing.T){
	s := []int{}
	//t.Log(s, len(s), cap(s))
	for i:=0; i < 10; i++ {
		s =append(s,i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T){
	year := []string{"Jan1", "Feb1", "Mar1", "Jan2", "Feb2", "Mar2", "Jan3", "Feb3", "Mar3","Jan4", "Feb4", "Mar4",}
	//t.Log(s, len(s), cap(s))
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))

	summer[0] = "Unkonw"
	t.Log(Q2, len(Q2), cap(Q2))
	t.Log(summer, len(summer), cap(summer))
	t.Log(year, len(year), cap(year))
}
