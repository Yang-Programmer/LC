package medium

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	nums := []int{2, 3, 4, 1, 4, 0, 4, -1, -2, -1}
	kFrequent := topKFrequent(nums, 2)
	fmt.Println(kFrequent)
	/*h := Heap{}
	h.Push(&EntryNums{
		Entry: 1,
		Nums:  3,
	})
	h.Push(&EntryNums{
		Entry: 2,
		Nums:  2,
	})
	h.Push(&EntryNums{
		Entry: 3,
		Nums:  2,
	})
	for _, v := range h.Arr {
		fmt.Println(v.Entry, v.Nums)
	}*/
}
