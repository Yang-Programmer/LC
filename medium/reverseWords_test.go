package medium

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	s := "  hello world  "
	words := reverseWords(s)
	fmt.Println(words)
}
