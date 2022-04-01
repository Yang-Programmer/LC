package easy

import "fmt"

func IsPalindrome(x int) bool {
	result := false
	if x < 0 || (x%10 == 0 && x!=0) {
		return result
	}
	rev := reverse(x)
    fmt.Println(rev)
	return rev==x
}

func reverse(x int) int {

	result := 0
	for  x !=0{
		result = result*10 + x%10
		x = x/10
	}
	return result
}
