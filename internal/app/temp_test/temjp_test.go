package temp_test

import (
	"fmt"
	"testing"
)

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
}

func TestTemp(t *testing.T){
	m := make(map[string]*student)
	m2 := make(map[string]student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for i, stu := range stus {
		temp := stu
		m[stu.Name] = &temp

		m2[stu.Name] = stus[i]
	}

	//t.Log(stus)
	//t.Log(&stus[1])
	//t.Log(m)
	t.Log(m["zhou"],m["li"],m["wang"],m["jj"])
	//t.Log(m2)
}

//1 兔子问题（斐波那契数列）
func sum(month int) int {
	fmt.Println("sum:", month)
	if month < 4 {
		return 2
	} else if month <= 0 {
		return 0
	}
	return sum(month-1) + sum(month-3)
}

//2 约梭芬杀人法
func kill(people []int) []int {
	if len(people) == 1 {
		return people
	}else if len(people) < 7 {
		a := 7%len(people)-1
		arr1 := append(people[a+1:],people[0:a]...)
		return kill(arr1)
	}
	arr1 := append(people[7:],people[0:6]...)
	return kill(arr1)
}
func TestKill(t *testing.T) {
	//fmt.Println("LLL")
	//1
	fmt.Println(sum(7))

	//2
	num := 12
	var arr [12]int
	for i := 0;i<num;i++{
		arr[i] = i+1
	}
	arr1 := arr[0:num]
	fmt.Println(arr)
	fmt.Println(kill(arr1))
}