package leetcode

var symbol = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) int {
	result := 0
	len := len(s)
	for i := 0; i < len; i++ {
		val := symbol[s[i]]
		if i < len-1 && val < symbol[s[i+1]] {
			result = result - val
		} else {
			result = result + val
		}
	}
	return result
}
