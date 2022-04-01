package easy

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T)  {
	nums := []int{3,2,2,3}
	val := 3
	n := removeElement(nums, val)
	for i:=0;i<n;i++{
		fmt.Print(nums[i],",")
	}
}
