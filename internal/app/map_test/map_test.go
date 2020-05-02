package map_test

import "testing"

func TestMapInit(t *testing.T)  {
	map1 := map[string]int {"aaa" :1,"bbb" :1,"ccc" :1}
	map1["ddd"] = 4
	t.Log(map1)

	m2 := map[int]int{}
	m2[4] = 16
	t.Log(m2, len(m2))

	m3 := make(map[int]int,2)
	t.Log(m3, len(m3))
}

func TestMapAccessNotExist(t *testing.T)  {
	m1 := map[int]int{}
	m1[2] = 0
	m1[3] = 0
	t.Log(m1[2])
	if v,ok := m1[3];ok{
		t.Logf("key 3 is %d", v)
	}else{
		t.Log("key 3 is not existing")
	}
}

func TestTravelMap(t *testing.T)  {
	m1 := map[string]int {"aaa" :1,"bbb" :1,"ccc" :1}
	for i,v := range m1 {
		t.Log(i,v)
	}
}
func TestMapWithFuncValue(t *testing.T){
	m := map[int]func(op int) int {}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op*op }
	m[3] = func(op int) int { return op*op*op }

	t.Log(m)
	t.Log(m[1](2),m[2](2),m[3](2))
}

func TestMapWithFuncValue2(t *testing.T){
	m := map[int]func(op int) (int, string) {}
	m[1] = func(op int) (int, string) { return op,"a" }
	m[2] = func(op int) (int, string) { return op*op,"aa" }
	m[3] = func(op int) (int, string) { return op*op*op,"aaa" }

	t.Log(m)

	x,y := m[3](2)
	t.Log(x,y)
	t.Log(m[1](2))
	t.Log(m[2](2))
	t.Log(m[3](2))
}