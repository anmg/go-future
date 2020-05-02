package main

import "../internal/app/square"

func main(){
	println("hello word")

	inputs := []int{1,2,3}
	expected := []int{1,4,9}
	for i :=0; i < len(inputs); i++ {
		ret := square.Square(inputs[i])
		if ret != expected[i]{
			println("%d, %d, %d", inputs[i], expected[i], ret)
		}
	}
}




