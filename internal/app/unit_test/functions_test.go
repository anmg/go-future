package functions_test

import (
	"testing"

)

func TestSquare(t *testing.T) {
	inputs := []int{1,2,3}
	expected := []int{1,4,9}
	for i :=0; i < len(inputs); i++ {
		ret := functions.Square(i)
		if ret != expected[i]{
			t.Logf("%d, %d, %d", inputs[i], expected[i], ret)
		}
	}
}