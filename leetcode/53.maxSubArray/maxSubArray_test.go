package leetcode

import (
	"fmt"
	"testing"
)

type question53 struct {
	param53
	answer53
}

type param53 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type answer53 struct {
	one int
}

func Test_Problem35(t *testing.T) {

	qs := []question53{
		{
			param53{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			answer53{6},
		},
		{
			param53{[]int{1}},
			answer53{0},
		},
		{
			param53{[]int{0}},
			answer53{0},
		},
		{
			param53{[]int{-1}},
			answer53{-1},
		},
		{
			param53{[]int{-100000}},
			answer53{-100000},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 53------------------------\n")

	for _, q := range qs {
		_, p := q.answer53, q.param53
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxSubArray(p.nums))
	}
	fmt.Printf("\n\n\n")
}
