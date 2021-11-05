package leetcode

import (
	"fmt"
	"testing"
)

type question217 struct {
	param217
	answer217
}

type param217 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type answer217 struct {
	one bool
}

func Test_Problem68(t *testing.T) {

	qs := []question217{
		{
			param217{[]int{1, 2, 3, 1}},
			answer217{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 217------------------------\n")

	for _, q := range qs {
		_, p := q.answer217, q.param217
		fmt.Printf("【input】:%v       【output】:%v\n", p, containsDuplicate(p.nums))
	}
	fmt.Printf("\n\n\n")
}
