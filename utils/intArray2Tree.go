package utils

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 完全二叉树的数组
func IntArray2Tree(arr []int) *TreeNode {
	l := len(arr)
	root := new(TreeNode)
	cur := root

	for i := 0; i < l; i++ {
		left, right := false, false
		if arr[i] == -1 {
			continue
		}
		lc := 2*i + 1
		rc := 2*i + 2
		cur.Val = arr[i]
		if lc < l && arr[lc] == -1 {
			cur.Left = nil
		}
		if lc < l && arr[lc] != -1 {
			cur.Left = &TreeNode{
				Val: arr[lc],
			}
			left = true
		}
		if rc < l && arr[rc] == -1 {
			cur.Right = nil
		}
		if rc < l && arr[rc] != -1 {
			cur.Right = &TreeNode{
				Val: arr[rc],
			}
			right = true
		}
		//如何移动cur指针
		if left && !right {
			cur = cur.Left
		}
		if left && right {
			cur = cur.Right
		}
		if !left && right {
			cur = cur.Right
		}
	}
	return root
}
