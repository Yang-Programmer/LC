package easy

import (
	"fmt"
	"testing"
)

func TestImplementStackUsingQueues(t *testing.T) {
	obj := ConstructorStack()
	obj.Push(1)
	obj.Push(2)
	param_2 := obj.Top()
	param_3 := obj.Pop()
	param_4 := obj.Empty()
	fmt.Println(param_2, param_3, param_4)
}
