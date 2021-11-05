package leetcode

import (
	"fmt"
	"testing"
)

type question58 struct {
	param58
	answer58
}

type param58 struct {
	nums string
}

// ans 是答案
// one 代表第一个答案
type answer58 struct {
	one int
}

func Test_Problem35(t *testing.T) {

	qs := []question58{
		{
			param58{"Hello World"},
			answer58{5},
		},
		{
			param58{"   fly me   to   the moon  "},
			answer58{4},
		},
		{
			param58{"luffy is still joyboy"},
			answer58{6},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 35------------------------\n")

	for _, q := range qs {
		_, p := q.answer58, q.param58
		fmt.Printf("【input】:%v       【output】:%v\n", p, lengthOfLastWord(p.nums))
	}
	fmt.Printf("\n\n\n")
}
