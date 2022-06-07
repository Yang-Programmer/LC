package medium

/**
堆排序 (升序)
*/

func heapSort(arr []int) {
	// 将混乱的初始数组堆化
	buildHeap(arr)
	//此时堆顶是最大值 将最大值与最后一个元素交换 并删除最后一个节点
	for i := len(arr) - 1; i >= 0; i-- {
		swap(arr, 0, i)
		heapify(arr, 0, i)
	}
}

func buildHeap(arr []int) {
	l := len(arr)
	parentIndex := (l - 1 - 1) / 2
	for ; parentIndex >= 0; parentIndex-- {
		heapify(arr, parentIndex, l)
	}
}
func heapify(arr []int, nodeIndex int, needLen int) {

	if nodeIndex >= needLen {
		return
	}
	leftChild := 2*nodeIndex + 1
	rightChild := 2*nodeIndex + 2
	max := nodeIndex
	if leftChild < needLen && arr[leftChild] > arr[max] {
		max = leftChild
	}
	if rightChild < needLen && arr[rightChild] > arr[max] {
		max = rightChild
	}
	if max != nodeIndex {
		swap(arr, nodeIndex, max)
		heapify(arr, max, needLen)
	}
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
