package hard

import (
	"fmt"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	ints := maxSlidingWindow(nums, 3)
	nums1 := []int{1}
	ints1 := maxSlidingWindow(nums1, 1)
	fmt.Println(ints, ints1)
}
