package easy

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(string(s))
	s1 := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	reverseString(s1)
	fmt.Println(string(s1))
}
