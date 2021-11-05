package leetcode

import (
	"fmt"
	"testing"
)

type question20 struct {
	param20
	answer20
}

type param20 struct {
	nums string
}

// ans 是答案
// one 代表第一个答案
type answer20 struct {
	one bool
}

func Test_Problem1(t *testing.T) {
	qs := []question20{
		{
			param20{"()[]{}"},
			answer20{true},
		},
		{
			param20{"()"},
			answer20{true},
		},
		{
			param20{"()["},
			answer20{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 20------------------------\n")
	for _, q := range qs {
		_, p := q.answer20, q.param20
		fmt.Printf("【input】:%v       【output】:%v\n", p, isValid(p.nums))
	}
	fmt.Printf("\n\n\n")
}
