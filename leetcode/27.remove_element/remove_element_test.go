package leetcode

import (
	"fmt"
	"testing"
)

type question27 struct {
	param27
	answer27
}

type param27 struct {
	nums []int
	val  int
}

// ans 是答案
// one 代表第一个答案
type answer27 struct {
	one int
}

func Test_Problem27(t *testing.T) {

	qs := []question27{
		{
			param27{[]int{1, 2, 3, 4, 4, 4, 4, 5, 6, 6, 7}, 4},
			answer27{7},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 27------------------------\n")

	for _, q := range qs {
		_, p := q.answer27, q.param27
		fmt.Printf("【input】:%v       【output】:%v\n", p, removeElement(p.nums, p.val))
	}
	fmt.Printf("\n\n\n")
}
