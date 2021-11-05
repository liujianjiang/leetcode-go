package leetcode

import (
	"fmt"
	"testing"
)

type question35 struct {
	param35
	answer35
}

type param35 struct {
	nums string
	val  string
}

// ans 是答案
// one 代表第一个答案
type answer35 struct {
	one int
}

func Test_Problem28(t *testing.T) {

	qs := []question35{
		{
			param35{"hello world", "ld"},
			answer35{9},
		},
		{
			param35{"aaaaa", "bba"},
			answer35{-1},
		},
		{
			param35{"aaaaa", ""},
			answer35{0},
		},
		{
			param35{"", ""},
			answer35{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 27------------------------\n")

	for _, q := range qs {
		_, p := q.answer35, q.param35
		fmt.Printf("【input】:%v       【output】:%v\n", p, strStr(p.nums, p.val))
	}
	fmt.Printf("\n\n\n")
}
