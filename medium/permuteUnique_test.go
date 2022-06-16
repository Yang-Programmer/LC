package medium

import (
	"fmt"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	unique := permuteUnique([]int{1, 1, 2})
	fmt.Println(unique)
}
