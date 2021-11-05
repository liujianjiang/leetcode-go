package leetcode

import (
	"fmt"
	"testing"
)

type question7 struct {
	param7
	answer7
}

type param7 struct {
	nums int
}

// ans 是答案
// one 代表第一个答案
type answer7 struct {
	one int
}

func Test_Problem1(t *testing.T) {

	qs := []question7{
		{
			param7{123456},
			answer7{654321},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")

	for _, q := range qs {
		_, p := q.answer7, q.param7
		fmt.Printf("【input】:%v       【output】:%v\n", p, reverse(p.nums))
	}
	fmt.Printf("\n\n\n")
}
