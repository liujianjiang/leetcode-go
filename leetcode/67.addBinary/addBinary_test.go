package leetcode

import (
	"fmt"
	"testing"
)

type question67 struct {
	param67
	answer67
}

type param67 struct {
	a string
	b string
}

// ans 是答案
// one 代表第一个答案
type answer67 struct {
	one string
}

func Test_Problem67(t *testing.T) {

	qs := []question67{
		{
			param67{"11", "1"},
			answer67{"100"},
		},
		{
			param67{"1010", "1011"},
			answer67{"10101"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 67------------------------\n")

	for _, q := range qs {
		_, p := q.answer67, q.param67
		fmt.Printf("【input】:%v       【output】:%v\n", p, addBinary(p.a, p.b))
	}
	fmt.Printf("\n\n\n")
}
