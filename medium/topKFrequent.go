package medium

/**
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。



示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]


提示：

1 <= nums.length <= 105
k 的取值范围是 [1, 数组中不相同的元素的个数]
题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/top-k-frequent-elements
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type EntryNums struct {
	Entry int
	Nums  int
}

type Heap struct {
	Arr []*EntryNums
}

func topKFrequent(nums []int, k int) []int {
	//先统计每个元素出现的次数
	showMap := make(map[int]*EntryNums)
	for _, entry := range nums {
		if _, ok := showMap[entry]; !ok {
			showMap[entry] = &EntryNums{
				Entry: entry,
			}
		}
		showMap[entry] = &EntryNums{
			Entry: entry,
			Nums:  showMap[entry].Nums + 1,
		}
	}
	h := new(Heap)
	h.Init()
	for _, numShow := range showMap {
		if h.Len() < k {
			h.Push(numShow)
		} else {
			if h.Arr[0].Nums < numShow.Nums {
				h.Pop()
				h.Push(numShow)
			}
		}
	}
	ret := make([]int, 0)
	for _, v := range h.Arr {
		ret = append(ret, v.Entry)
	}
	return ret
}

func (h *Heap) MHeapify(position int, length int) {
	if position >= length {
		return
	}
	leftChild := 2*position + 1
	rightChild := 2*position + 2
	min := position
	if leftChild < length && h.Arr[leftChild].Nums < h.Arr[min].Nums {
		min = leftChild
	}
	if rightChild < length && h.Arr[rightChild].Nums < h.Arr[min].Nums {
		min = rightChild
	}
	if min != position {
		h.Swap(min, position)
		h.MHeapify(min, length)
	}
}

func (h *Heap) Init() {
	h.Arr = make([]*EntryNums, 0)
}
func (h *Heap) Len() int {
	return len(h.Arr)
}
func (h *Heap) Swap(i, j int) {
	tmp := h.Arr[i]
	h.Arr[i] = h.Arr[j]
	h.Arr[j] = tmp
}
func (h *Heap) Push(x *EntryNums) {
	h.Arr = append(h.Arr, x)
	for i := h.Len()/2 - 1; i >= 0; i-- {
		h.MHeapify(i, h.Len())
	}
}

func (h *Heap) Pop() {
	h.Swap(0, h.Len()-1)
	h.Arr = h.Arr[:h.Len()-1]
}
