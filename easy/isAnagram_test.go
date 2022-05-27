package easy

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	s := "anagram"
	s2 := "nagaram"
	s3 := "rat"
	s4 := "car"
	result := isAnagram(s, s2)
	result2 := isAnagram(s3, s4)
	fmt.Println(result, result2)
}
