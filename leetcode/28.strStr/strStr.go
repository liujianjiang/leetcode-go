package leetcode

import "fmt"

func strStr(haystack string, needle string) int {

	hlen := len(haystack)
	nlen := len(needle)

	if hlen == 0 && nlen == 0 {
		return 0
	}
	if hlen != 0 && nlen == 0 {
		return 0
	}

	if nlen == 0 || nlen > hlen {
		return -1
	}

	head := 0
	tail := nlen

	num := -1
	for i := 0; i < hlen-nlen+1; i++ {
		fmt.Println(haystack[head:tail])
		if haystack[head:tail] == needle {
			num = head
			break
		}
		head++
		tail++
	}
	return num
}
