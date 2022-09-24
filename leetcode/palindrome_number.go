package leetcode

import "fmt"

func isPalindrome(x int) bool {
	x_ := fmt.Sprint(x)
	r := []rune(x_)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return x_ == string(r)
}
