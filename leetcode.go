package main

import (
	"fmt"
	"leetcode/medium"
)

func main() {
	/*nums := []int{3, 2, 4}
	var target int = 6
	re := easy.TwoSum(nums, target)
	fmt.Println(re)*/

	/*var nums = -123
	re := easy.ReverseV2(nums)
	fmt.Println(re)*/

	/*num := 10
	re:=easy.IsPalindrome(num)
	fmt.Println(re)*/

		var (
			l1 = new(medium.ListNode)
			l2 = new(medium.ListNode)
			l3 = new(medium.ListNode)
			l4 = new(medium.ListNode)
			l5 = new(medium.ListNode)
			l6 = new(medium.ListNode)
			/*l7 = new(medium.ListNode)

			l8  = new(medium.ListNode)
			l9  = new(medium.ListNode)
			l10 = new(medium.ListNode)
			l11 = new(medium.ListNode)
			l12 = new(medium.ListNode)
			l13 = new(medium.ListNode)
			l14 = new(medium.ListNode)
			l15 = new(medium.ListNode)
			l16 = new(medium.ListNode)
			l17 = new(medium.ListNode)
			l18 = new(medium.ListNode)
			l19 = new(medium.ListNode)
			l20 = new(medium.ListNode)
			l21 = new(medium.ListNode)
			l22 = new(medium.ListNode)
			l23 = new(medium.ListNode)
			l24 = new(medium.ListNode)
			l25 = new(medium.ListNode)
			l26 = new(medium.ListNode)
			l27 = new(medium.ListNode)
			l28 = new(medium.ListNode)
			l29 = new(medium.ListNode)
			l30 = new(medium.ListNode)
			l31 = new(medium.ListNode)
			l32 = new(medium.ListNode)
			l33 = new(medium.ListNode)
			l34 = new(medium.ListNode)
			l35 = new(medium.ListNode)
			l36 = new(medium.ListNode)
			l37 = new(medium.ListNode)
			l38 = new(medium.ListNode)*/
		)
		l1.Next = l2
		l2.Next = l3
		l1.Val = 2
		l2.Val = 4
		l3.Val = 3

		//l7.Val = 0

		l4.Next = l5
		l5.Next = l6
		l4.Val = 5
		l5.Val = 6
		l6.Val = 4

	/*	l8.Next = l9
		l9.Next = l10
		l10.Next = l11
		l11.Next = l12
		l12.Next = l13
		l13.Next = l14
		l14.Next = l15
		l15.Next = l16
		l16.Next = l17
		l17.Next = l18
		l18.Next = l19
		l19.Next = l20
		l20.Next = l21
		l21.Next = l22
		l22.Next = l23
		l23.Next = l24
		l24.Next = l25
		l25.Next = l26
		l26.Next = l27
		l27.Next = l28
		l28.Next = l29
		l29.Next = l30
		l30.Next = l31
		l31.Next = l32
		l32.Next = l33
		l33.Next = l34
		l34.Next = l35
		l35.Next = l36
		l36.Next = l37
		l37.Next = l38

		l8.Val = 1
		l38.Val = 1*/

		//l := medium.AddTwoNumbers(l1, l4)
		l := medium.AddTwoNumbersV2(l1, l4)
		for l != nil {
			fmt.Println(l.Val)
			l = l.Next
		}

	/*s := "bbbbb"
	longestSubstring := medium.LengthOfLongestSubstring(s)
	fmt.Println(longestSubstring)*/
}
