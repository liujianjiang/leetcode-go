package leetcode

import (
	"fmt"
	"testing"
)

type question66 struct {
	param66
	answer66
}

type param66 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type answer66 struct {
	one int
}

func Test_Problem35(t *testing.T) {

	qs := []question66{
		{
			param66{[]int{1, 2, 3}},
			answer66{123},
		},
		{
			param66{[]int{4, 3, 2, 1}},
			answer66{4321},
		},
		{
			param66{[]int{0}},
			answer66{1},
		},
		{
			param66{[]int{9, 9}},
			answer66{100},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 66------------------------\n")

	for _, q := range qs {
		_, p := q.answer66, q.param66
		fmt.Printf("【input】:%v       【output】:%v\n", p, plusOne(p.nums))
	}
	fmt.Printf("\n\n\n")
}
