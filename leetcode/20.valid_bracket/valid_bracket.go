package leetcode

import "fmt"

func isValid(s string) bool {
	length := len(s)
	if length%2 == 1 {
		return false
	}
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := []byte{}
	for i := 0; i < length; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			} else {
				fmt.Printf("find : %s\n", string(s[i]))
			}
			//fmt.Printf("chu : %s \n", stack)
			stack = stack[:len(stack)-1]
		} else {
			fmt.Printf("ru : %s\n", string(s[i]))
			stack = append(stack, s[i])
		}
		//fmt.Printf("chu : %s \n", stack)
	}

	return len(stack) == 0
}
