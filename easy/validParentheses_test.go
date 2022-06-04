package easy

import (
	"fmt"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	s := "()"
	s1 := "()[]{}"
	s2 := "(]"
	s3 := "([)]"
	s4 := "{[]}"
	valid := isValid(s)
	valid1 := isValid(s1)
	valid2 := isValid(s2)
	valid3 := isValid(s3)
	valid4 := isValid(s4)
	fmt.Println(valid, valid1, valid2, valid3, valid4)
}
