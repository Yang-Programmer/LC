package medium

import (
	"fmt"
	"testing"
)

func TestJump(t *testing.T) {
	step := jump([]int{2, 1, 1, 1, 4})
	fmt.Println(step)
}
