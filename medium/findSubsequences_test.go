package medium

import (
	"fmt"
	"testing"
)

func TestFindSubsequences(t *testing.T) {
	subsequences := findSubsequences([]int{4, 6, 7, 7})
	fmt.Println(subsequences)
}
