package easy

import (
	"fmt"
	"testing"
)

func TestReverseStrV2(t *testing.T) {
	str := reverseStr("abcdefg", 2)
	str2 := reverseStr("abcd", 2)
	fmt.Println(str, str2)
}
