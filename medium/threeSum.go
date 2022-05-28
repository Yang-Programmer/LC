package medium

import "sort"

/**
给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
示例 2：

输入：nums = []
输出：[]
示例 3：

输入：nums = [0]
输出：[]


提示：

0 <= nums.length <= 3000
-105 <= nums[i] <= 105

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	length := len(nums)
	for i := 0; i < length-2; i++ {
		now := nums[i]
		if now > 0 {
			break
		}
		if i > 0 && nums[i-1] == now {
			continue
		}
		left, right := i+1, length-1
		for left < right {
			n1, n2 := nums[left], nums[right]
			if now+n1+n2 == 0 {
				result = append(result, []int{now, n1, n2})
				for left < right && nums[left] == n1 {
					left++
				}
				for left < right && nums[right] == n2 {
					right--
				}
			} else if now+n1+n2 < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
