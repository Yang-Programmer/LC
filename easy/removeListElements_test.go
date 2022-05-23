package easy

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestRemoveListElements(t *testing.T) {
	head := []int{1, 2, 6, 3, 4, 5, 6}
	listNodes := utils.Array2ListInt(head)
	elements := removeElements(listNodes, 6)
	for elements != nil {
		fmt.Println(elements.Val)
		elements = elements.Next
	}
}
