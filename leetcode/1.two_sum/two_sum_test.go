package leetcode

import (
	"fmt"
	"testing"
)

type question1 struct {
	param
	answer
}

type param struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type answer struct {
	one []int
}

func Test_Problem1(t *testing.T) {

	qs := []question1{
		{
			param{[]int{1, 2, 3, 4}, 7},
			answer{[]int{2, 3}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")

	for _, q := range qs {
		_, p := q.answer, q.param
		fmt.Printf("【input】:%v       【output】:%v\n", p, twoSum(p.nums, p.target))
	}
	fmt.Printf("\n\n\n")
}
