package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "111", 2: "222", 3: "333"}
	b := map[int]string{1: "111", 2: "222", 3: "333"}

	//fmt.Println(a == b)
	fmt.Println(reflect.DeepEqual(a, b))

	s1 := []int{1,2,3}
	s2 := []int{1,2,3}
	s3 := []int{1,2,3}
	t.Log("s1 s2", reflect.DeepEqual(s1, s2))
	t.Log("s1 s3", reflect.DeepEqual(s1, s3))
}

type Employee struct {
	EmployeeID string
	Name string
	Age int
}


type Customer struct {
	CustomerID string
	Name string
	Age int
}

type MyInt int

func TestReflect( t *testing.T){
	//var x float64 = 3.4
	//b := reflect.ValueOf(x)
	//
	//fmt.Println(b.Type(), b.Kind(), b.Float(), b.Int())
	//fmt.Println("type:", reflect.TypeOf(x))
	//fmt.Println("value:", reflect.ValueOf(x))

	//var x float64 = 7.1
	//
	//v := reflect.ValueOf(x)
	//fmt.Println(v.Type(), v.Kind())


	//y := v.Interface().(MyInt) // y will have type float64.

	//fmt.Println(v.Interface())
	//
	//fmt.Printf("value is %7.1e\n", v.Interface())
	//fmt.Println(v)

	//var x1 float64 = 3.4
	//v1 := reflect.ValueOf(x1)
	//var x2 float64 = 7.1
	//v1.SetFloat(x2) // Error: will panic.


	//var x float64 = 3.4
	//v := reflect.ValueOf(&x)
	//v.SetFloat(7.1)
	//fmt.Println(v.CanSet())


	var x float64 = 3.4
	p := reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())


	v := p.Elem()
	fmt.Println("settability of v:", v.CanSet())

	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)
}

type T1 struct {
	A int
	B string
}

//反射从接口值到反射对象中(Reflection goes from interface value to reflection object.)
//反射从反射对象到接口值中(Reflection goes from reflection object to interface value.)
//要修改反射对象，值必须是“可设置”的(To modify a reflection object, the value must be settable.)
func TestReflectStruct( t *testing.T){
	t1 := T1{23, "skidoo"}
	s := reflect.ValueOf(&t1).Elem()
	typeOfT := s.Type()

	fmt.Println(typeOfT)
	for i := 0; i < s.NumField(); i++ {
		//fmt.Println(i, s.Field(i))
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v %t\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface(), f.CanSet())

	}

	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t1, "t is now", t1,"t is now", t1)

}