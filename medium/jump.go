package medium

/**
给你一个非负整数数组nums ，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。

假设你总是可以到达数组的最后一个位置。



示例 1:

输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
    从下标为 0 跳到下标为 1 的位置，跳1步，然后跳3步到达数组的最后一个位置。
示例 2:

输入: nums = [2,3,0,1,4]
输出: 2


提示:

1 <= nums.length <= 104
0 <= nums[i] <= 1000

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/jump-game-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func jump(nums []int) int {
	count := 0
	currentIndex, nextIndex := 0, 0
	l := len(nums)
	if l == 1 {
		return count
	}
	for i := 0; i < l-1; i++ {
		if nums[i]+i > nextIndex {
			nextIndex = nums[i] + i
		}
		if currentIndex == i {
			count++
			currentIndex = nextIndex
		}
	}
	return count
}
