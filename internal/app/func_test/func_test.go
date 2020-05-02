package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValus() (int, int){
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)

		fmt.Println("time spent:", time.Since(start).Seconds())

		return ret
	}

}

func slowFunc(op int) int {
	time.Sleep(time.Second*1)
	return op
}

func Sum(ops...int) int {
	ret := 0
	for _,op := range ops {
		ret += op
	}
	return ret
}

func TestFunc(t *testing.T){
	a,b := returnMultiValus()

	t.Log(a,b)

	tsSF := timeSpent(slowFunc)
	t.Log(tsSF(10))

	t.Log(Sum(1,2,3,4))
	t.Log(Sum(1,2,3,4,5))
}

func Clear(){
	fmt.Println("------------------------clear resources")
}

func TestDefer(t *testing.T){
	defer Clear()
	fmt.Println("-------------------------start")
	panic("Fatal error")
	fmt.Println("-------------------------end")
}

//func TestDefer2(t *testing.T){
//	defer func() {
//		t.Log("clear")
//	}()
//	t.Log("start 2")
//	//panic("Fatal error")
//}

