package easy

import (
	"fmt"
	"testing"
)

func TestReverseLeftWords(t *testing.T) {
	s := "abcdefg"
	words := reverseLeftWords(s, 2)
	s1 := "lrloseumgh"
	words1 := reverseLeftWords(s1, 6)
	fmt.Println(words, words1)
}
