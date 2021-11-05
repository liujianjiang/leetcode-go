package leetcode

import (
	"fmt"
	"testing"
)

type question70 struct {
	param70
	answer70
}

type param70 struct {
	nums int
}

// ans 是答案
// one 代表第一个答案
type answer70 struct {
	one int
}

func Test_Problem68(t *testing.T) {

	qs := []question70{
		{
			param70{2},
			answer70{2},
		},
		{
			param70{3},
			answer70{3},
		},
		{
			param70{5},
			answer70{8},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 70------------------------\n")

	for _, q := range qs {
		_, p := q.answer70, q.param70
		fmt.Printf("【input】:%v       【output】:%v\n", p, climbStairs(p.nums))
	}
	fmt.Printf("\n\n\n")
}
