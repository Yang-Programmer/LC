package easy

import (
	"container/list"
	"leetcode/utils"
)

/**
给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。



示例 1：


输入：root = [1,null,2,3]
输出：[3,2,1]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]


提示：

树中节点的数目在范围 [0, 100] 内
-100 <= Node.val <= 100


进阶：递归算法很简单，你可以通过迭代算法完成吗？

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/binary-tree-postorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// 迭代法（使用栈）实现二叉树的后序迭代
/**
后序遍历:
先序遍历是中左右，后续遍历是左右中，
那么我们只需要调整一下先序遍历的代码顺序，就变成中右左的遍历顺序，
然后在反转result数组，输出的结果顺序就是左右中了
*/
func postorderTraversal(root *utils.TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		top := stack.Remove(stack.Back()).(*utils.TreeNode)
		ret = append(ret, top.Val)
		if top.Left != nil {
			stack.PushBack(top.Left)
		}
		if top.Right != nil {
			stack.PushBack(top.Right)
		}
	}
	reverserArr := func(arr []int) {
		i, j := 0, len(arr)-1
		for i < j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	reverserArr(ret)
	return ret
}
