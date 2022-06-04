package medium

import (
	"fmt"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	tokens := []string{"2", "1", "+", "3", "*"}
	tokens1 := []string{"4", "13", "5", "/", "+"}
	tokens2 := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	rpn := evalRPN(tokens)
	rpn1 := evalRPN(tokens1)
	rpn2 := evalRPN(tokens2)
	fmt.Println(rpn, rpn1, rpn2)
}
