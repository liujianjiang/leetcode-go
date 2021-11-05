package leetcode

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if prefix == "" {
			break
		}
	}
	return prefix
}

func lcp(str1 string, str2 string) string {
	length := 0
	if len(str1) < len(str2) {
		length = len(str1)
	} else {
		length = len(str2)
	}
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}
