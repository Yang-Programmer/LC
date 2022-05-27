package easy

import (
	"fmt"
	"testing"
)

func TestIntersection(t *testing.T) {
	ar1, ar2 := []int{1, 2, 2, 1}, []int{2, 2}
	ar3, ar4 := []int{4, 9, 5}, []int{9, 4, 9, 8, 4}
	result := intersection(ar1, ar2)
	result1 := intersection(ar3, ar4)
	fmt.Println(result, result1)
}
