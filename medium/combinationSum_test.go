package medium

import (
	"fmt"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	ret := combinationSum([]int{2, 3, 6, 7}, 7)
	fmt.Println(ret)
}
