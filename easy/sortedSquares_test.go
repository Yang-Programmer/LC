package easy

import (
	"fmt"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	nums := []int{-7, -3, 2, 3, 11}
	squares := sortedSquares(nums)
	fmt.Println(squares)
}
