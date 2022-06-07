package medium

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{1, 8, 2, 4, 9, 10}
	heapSort(arr)
	for _, entry := range arr {
		fmt.Println(entry)
	}
}
