package easy

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	arr := []int{1, -1, 2, -1, -1, 3}
	tree := utils.IntArray2Tree(arr)
	ints := inorderTraversal(tree)
	fmt.Println(ints)

	arr1 := []int{}
	tree1 := utils.IntArray2Tree(arr1)
	ints1 := inorderTraversal(tree1)
	fmt.Println(ints1)

	arr2 := []int{1}
	tree2 := utils.IntArray2Tree(arr2)
	ints2 := inorderTraversal(tree2)
	fmt.Println(ints2)
}

func TestStackInorderTraversal(t *testing.T) {
	arr := []int{1, -1, 2, -1, -1, 3}
	tree := utils.IntArray2Tree(arr)
	ints := inorderTraversalV2(tree)
	fmt.Println(ints)
}
