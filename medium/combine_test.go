package medium

import (
	"fmt"
	"testing"
)

func TestCombine(t *testing.T) {
	ret := combine(4, 2)
	ret1 := combineV2(4, 2)
	fmt.Println(ret, ret1)
}
