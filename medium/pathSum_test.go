package medium

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestPathSum(t *testing.T) {
	root := &utils.TreeNode{
		Val: 5,
		Left: &utils.TreeNode{
			Val: 4,
			Left: &utils.TreeNode{
				Val: 11,
				Left: &utils.TreeNode{
					Val: 7,
				},
				Right: &utils.TreeNode{
					Val: 2,
				},
			},
		},
		Right: &utils.TreeNode{
			Val: 8,
			Left: &utils.TreeNode{
				Val: 13,
			},
			Right: &utils.TreeNode{
				Val: 4,
				Left: &utils.TreeNode{
					Val: 5,
				},
				Right: &utils.TreeNode{
					Val: 1,
				},
			},
		},
	}
	ret := pathSum(root, 22)
	fmt.Println(ret)
}
