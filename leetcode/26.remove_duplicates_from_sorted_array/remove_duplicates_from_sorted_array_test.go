package leetcode

import (
	"fmt"
	"testing"
)

type question26 struct {
	param26
	answer26
}

type param26 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type answer26 struct {
	one int
}

func Test_Problem26(t *testing.T) {

	qs := []question26{
		{
			param26{[]int{1, 1, 2}},
			answer26{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 26------------------------\n")

	for _, q := range qs {
		_, p := q.answer26, q.param26
		fmt.Printf("【input】:%v       【output】:%v\n", p, removeDuplicates(p.nums))
	}
	fmt.Printf("\n\n\n")
}
