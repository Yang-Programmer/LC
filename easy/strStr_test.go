package easy

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	idx := strStr("hello", "ll")
	idx1 := strStr("aaaaa", "bba")
	idx2 := strStr("mississippi", "mississippi")
	fmt.Println("result: ", idx, idx1, idx2)
}
