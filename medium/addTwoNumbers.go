package medium

import (
	"fmt"
	"math"
	"strconv"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/* 错误版本
错误原因：当数量过于大的时候 float64会产生位数溢出导致错误答案
*/
func AddTwoNumbersWrong(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	list := new(ListNode)
	var (
		rL1 float64 = 0
		rL2 float64 = 0
	)

	var i float64 = 0
	for l1 != nil {
		if l1.Val > 0 {
			rL1 = rL1 + float64(l1.Val)*math.Pow(10, i)
		} else {
			rL1 = rL1 * 10
		}

		i = i + 1
		str := strconv.FormatFloat(rL1, 'f', -1, 128)
		fmt.Println(l1.Val, str)
		l1 = l1.Next
	}

	var j float64 = 0
	for l2 != nil {
		rL2 = rL2 + float64(l2.Val)*math.Pow(10, j)
		j = j + 1
		l2 = l2.Next
	}
	vResult := rL1 + rL2

	var k = 0
	for vResult != 0 {
		tmp := new(ListNode)
		tmp.Val = int(vResult) % 10
		vResult = float64(int(vResult) / 10)
		list.Next = tmp
		if k == 0 {
			result = tmp
		}
		list = list.Next
		k++
	}
	return result
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		result *ListNode
		list   *ListNode
	)

	carry := 0
	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2 + carry
		sum, carry = sum%10, sum/10
		if result == nil {
			result = &ListNode{
				Val: sum,
			}
			list = result
		} else {
			list.Next = &ListNode{
				Val: sum,
			}
			list = list.Next
		}
	}
	if carry > 0 {
		list.Next = &ListNode{
			Val: carry,
		}
	}
	return result
}

/* 错误版本
错误原因：当数量过于大的时候 float64会产生位数溢出导致错误答案
再次犯错
*/
func AddTwoNumbersV2(l1, l2 *ListNode) *ListNode {
	var sumL1 int
	var sumL2 int
	resultHeader := new(ListNode)

	times := 0
	for l1 != nil {
		tmp := l1.Val * pow(10, times)
		sumL1 += tmp
		times++
		l1 = l1.Next
	}
	times = 0
	for l2 != nil {
		tmp := l2.Val * pow(10, times)
		sumL2 += tmp
		times++
		l2 = l2.Next
	}
	sum := sumL1 + sumL2
	if sum == 0 {
		return &ListNode{
			Val: 0,
		}
	}
	resultTail := new(ListNode)
	resultHeader = resultTail
	for sum != 0 {
		tmp := &ListNode{
			Val: sum % 10,
		}
		resultTail.Next = tmp
		resultTail = resultTail.Next
		sum = sum / 10
	}

	return resultHeader.Next
}

func pow(a, b int) int {
	var result int = a
	if b == 0 {
		result = 1
	}
	for i := 0; i < b-1; i++ {
		result *= a
	}
	return result
}
