package medium

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	inputArr := []int{-1, 0, 1, 2, -1, -4}
	inputArr1 := []int{}
	inputArr2 := []int{0}
	result := threeSum(inputArr)
	result1 := threeSum(inputArr1)
	result2 := threeSum(inputArr2)
	fmt.Println(result, result1, result2)
}
