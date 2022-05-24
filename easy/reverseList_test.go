package easy

import (
	"leetcode/utils"
	"testing"
)

func TestReverseList(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	linkList := utils.Array2ListInt(arr)
	reverse := reverseList(linkList)

	utils.PrintLinkListElements(reverse)
}
