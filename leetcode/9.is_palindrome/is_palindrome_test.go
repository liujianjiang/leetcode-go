package leetcode

import (
	"fmt"
	"testing"
)

type question9 struct {
	param9
	answer9
}

type param9 struct {
	nums int
}

// ans 是答案
// one 代表第一个答案
type answer9 struct {
	one bool
}

func Test_Problem1(t *testing.T) {

	qs := []question9{
		{
			param9{12321},
			answer9{true},
		},
		{
			param9{123321},
			answer9{true},
		},
		{
			param9{123123},
			answer9{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")

	for _, q := range qs {
		_, p := q.answer9, q.param9
		fmt.Printf("【input】:%v       【output】:%v\n", p, isPalindrome(p.nums))
	}
	fmt.Printf("\n\n\n")
}
