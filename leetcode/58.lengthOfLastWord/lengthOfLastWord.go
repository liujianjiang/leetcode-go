package leetcode

func lengthOfLastWord(s string) int {
	n := len(s)

	tmp := ""

	for i := n - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if len(tmp) > 0 {
				break
			}
		} else {
			tmp += string(s[i])
		}
	}
	return len(tmp)
}
