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
	one int
}

func Test_Problem1(t *testing.T) {

	qs := []question20{
		{
			param20{"III"},
			answer20{3},
		},
		{
			param20{"IV"},
			answer20{4},
		},
		{
			param20{"IX"},
			answer20{9},
		},
		{
			param20{"LVIII"},
			answer20{58},
		},
		{
			param20{"MCMXCIV"},
			answer20{1994},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")
	for _, q := range qs {
		_, p := q.answer20, q.param20
		fmt.Printf("【input】:%v       【output】:%v\n", p, romanToInt(p.nums))
	}
	fmt.Printf("\n\n\n")
}
