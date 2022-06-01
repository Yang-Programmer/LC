package medium

import (
	"fmt"
	"testing"
)

func TestGetKmpNext(t *testing.T) {
	next := getNext("ABCAACBBCBADAABCACBD")
	fmt.Println(next)
}

func TestKmpStr(t *testing.T) {
	idx := strStr("hello", "ll")
	idx1 := strStr("aaaaa", "bba")
	idx2 := strStr("mississippi", "mississippi")
	fmt.Println("result: ", idx, idx1, idx2)
}
