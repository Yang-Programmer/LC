package easy

import (
	"math"
)

// IntMax is max int
//const IntMax = int(^uint(0) >> 1)

// IntMin is min int
//const IntMin = ^IntMax

// Reverse 整数反转
func Reverse(x int) int {
	var (
		result int = 0
		tmp    int = 0
	)
	max := math.Pow(2,31)-1
	min := 0-math.Pow(2,31)
	if float64(x) > max || float64(x) < min {
		return 0
	}
	if x < 0 {
		x = 0 - x
		for x != 0 {
			tmp = tmp*10 + x%10
			x = x / 10
		}
		result = 0 - tmp
	} else {
		for x != 0 {
			tmp = tmp*10 + x%10
			x = x / 10
		}
		result = tmp
	}
	if float64(result) > max || float64(result) < min {
		return 0
	}
	return result
}

// Reverse 整数反转
func ReverseV2(x int) int {
	var (
		result int = 0
		tmp    int = 0
	)
	max := math.Pow(2,31)-1
	min := 0-math.Pow(2,31)
	if float64(x) > max || float64(x) < min {
		return 0
	}
	for x != 0 {
		tmp = tmp*10 + x%10
		x = x / 10
	}
	result = tmp
	if float64(result) > max || float64(result) < min {
		return 0
	}
	return result
}
