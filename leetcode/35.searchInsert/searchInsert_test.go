package leetcode

import (
	"fmt"
	"testing"
)

type question35 struct {
	param35
	answer35
}

type param35 struct {
	nums []int
	val  int
}

// ans 是答案
// one 代表第一个答案
type answer35 struct {
	one int
}

func Test_Problem35(t *testing.T) {

	qs := []question35{
		{
			param35{[]int{1, 3, 5, 6}, 5},
			answer35{2},
		},
		{
			param35{[]int{1, 3, 5, 6}, 2},
			answer35{1},
		},
		{
			param35{[]int{1, 3, 5, 6}, 7},
			answer35{4},
		},
		{
			param35{[]int{1, 3, 5, 6}, 0},
			answer35{0},
		},
		{
			param35{[]int{1}, 0},
			answer35{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 35------------------------\n")

	for _, q := range qs {
		_, p := q.answer35, q.param35
		fmt.Printf("【input】:%v       【output】:%v\n", p, searchInsert(p.nums, p.val))
	}
	fmt.Printf("\n\n\n")
}
