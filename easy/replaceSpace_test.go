package easy

import (
	"fmt"
	"testing"
)

func TestReplaceSpace(t *testing.T) {
	s := "We are happy."
	replaceStr := replaceSpace(s)
	fmt.Println(replaceStr)
}
