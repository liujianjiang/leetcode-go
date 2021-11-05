package leetcode

import (
	"math"
)

func reverse(x int) int {
	var rev = 0
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x = x / 10
		rev = rev*10 + digit
	}
	return rev
}
