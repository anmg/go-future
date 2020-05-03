package empty_interface

import (
	"fmt"
	"testing"
)

var aaa struct{

}

func DoSomething(p interface{}) {
	//if i, ok := p.(int); ok {
	//	fmt.Println("Integer ",i)
	//	return
	//}
	//if i, ok := p.(string); ok {
	//	fmt.Println("String ",i)
	//	return
	//}
	//
	//fmt.Println("Unknow ")

	if i, ok := p.(int); ok {
		fmt.Println("Integer ",i)
		return
	}else if i, ok := p.(string); ok {
		fmt.Println("String ",i)
		return
	}else {
		fmt.Println("Unknow ")
	}
	//switch v:= p.(type) {
	//case int:
	//	fmt.Println("Integer ",v)
	//case string:
	//	fmt.Println("String ",v)
	//default:
	//	fmt.Println("Unknow ")
	//}
}

func TestEmptyInterface(t *testing.T)  {
	DoSomething(10)
	DoSomething("44")
	DoSomething(3.24)

}
