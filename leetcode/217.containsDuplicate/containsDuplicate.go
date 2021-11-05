package leetcode

func containsDuplicate(nums []int) bool {
	maps := map[int]struct{}{}

	for _, v := range nums {
		if _, has := maps[v]; has {
			return true
		}
		maps[v] = struct{}{}
	}
	return false
}
