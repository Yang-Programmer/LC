package easy

import (
	"math"
	"sort"
)

/**
给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：

选择某个下标 i并将 nums[i] 替换为 -nums[i] 。
重复这个过程恰好 k 次。可以多次选择同一个下标 i 。

以这种方式修改数组后，返回数组 可能的最大和 。



示例 1：

输入：nums = [4,2,3], k = 1
输出：5
解释：选择下标 1 ，nums 变为 [4,-2,3] 。
示例 2：

输入：nums = [3,-1,0,2], k = 3
输出：6
解释：选择下标 (1, 2, 2) ，nums 变为 [3,1,0,2] 。
示例 3：

输入：nums = [2,-3,-1,5,-4], k = 2
输出：13
解释：选择下标 (1, 4) ，nums 变为 [2,3,-1,5,4] 。


提示：

1 <= nums.length <= 104
-100 <= nums[i] <= 100
1 <= k <= 104


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximize-sum-of-array-after-k-negations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func largestSumAfterKNegations(nums []int, k int) int {
	sum := 0
	l := len(nums)
	//sortNums(nums)
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})
	for i := 0; i < l && k > 0; i++ {
		if nums[i] < 0 {
			nums[i] = 0 - nums[i]
			k--
		}
	}
	if k > 0 {
		for i := 0; i < k; i++ {
			nums[l-1] = 0 - nums[l-1]
			k--
		}
	}
	for i := 0; i < l; i++ {
		sum += nums[i]
	}
	return sum
}

func sortNums(nums []int) {
	partition := func(nums []int, left, right int) (middle int) {
		middle = left
		for idx := left + 1; idx <= right; idx++ {
			if absInt(nums[idx]) > absInt(nums[middle]) {
				tmp := nums[idx]
				nums[idx] = nums[middle]
				nums[middle] = tmp
			}
		}
		return
	}
	var quickSort func(nums []int, left, right int)
	quickSort = func(nums []int, left, right int) {
		if left > right {
			return
		}
		middle := partition(nums, left, right)
		quickSort(nums, left, middle-1)
		quickSort(nums, middle+1, right)
	}
	quickSort(nums, 0, len(nums)-1)
	return
}
func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
