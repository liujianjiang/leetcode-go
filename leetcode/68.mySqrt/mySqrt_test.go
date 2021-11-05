package leetcode

import (
	"fmt"
	"testing"
)

type question68 struct {
	param68
	answer68
}

type param68 struct {
	nums int
}

// ans 是答案
// one 代表第一个答案
type answer68 struct {
	one int
}

func Test_Problem68(t *testing.T) {

	qs := []question68{
		{
			param68{4},
			answer68{2},
		},
		{
			param68{8},
			answer68{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 68------------------------\n")

	for _, q := range qs {
		_, p := q.answer68, q.param68
		fmt.Printf("【input】:%v       【output】:%v\n", p, mySqrt(p.nums))
	}
	fmt.Printf("\n\n\n")
}
