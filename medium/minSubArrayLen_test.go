package medium

import (
	"fmt"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	subArrayLen := minSubArrayLen(target, nums)
	fmt.Println(subArrayLen)

	nums2 := []int{1,4,4}
	target2 := 4
	subArrayLen2 := minSubArrayLen(target2, nums2)
	fmt.Println(subArrayLen2)

	nums3 := []int{1,1,1,1,1,1,1,1}
	target3 := 11
	subArrayLen3 := minSubArrayLen(target3, nums3)
	fmt.Println(subArrayLen3)

	nums4 := []int{1,2,3,4,5}
	target4 := 15
	subArrayLen4 := minSubArrayLen(target4, nums4)
	fmt.Println(subArrayLen4)
}
