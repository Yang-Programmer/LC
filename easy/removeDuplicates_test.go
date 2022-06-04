package easy

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	s := "abbaca"
	ss := removeDuplicates(s)
	fmt.Println(ss)
}
