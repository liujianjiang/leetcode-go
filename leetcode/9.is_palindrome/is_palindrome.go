package leetcode

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x = x / 10
		fmt.Println(x)
	}
	//fmt.Println(revertedNumber)
	return x == revertedNumber || x == revertedNumber/10
}
