package map_ext_test

import "testing"

func TestMapWithFunValue(t *testing.T){
	m := map[int]func(op int) int {}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op*op }
	m[3] = func(op int) int { return op*op*op }

	t.Log(m)
	t.Log(m[1](2),m[2](2),m[3](2))
}

func TestMapForSet(t *testing.T){
	myset := map[int] bool{}
	myset[1] = true
	myset[2] = false
	n := 1
	t.Log(myset)
	if _,ok := myset[n];ok{
		t.Log("1 is exist")
	}else{
		t.Log("1 is not exist")
	}

	delete(myset, 1)
	t.Log(myset)
}