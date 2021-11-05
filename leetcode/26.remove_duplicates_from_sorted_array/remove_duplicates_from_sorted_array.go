package leetcode

func removeDuplicates(nums []int) int {
	lenght := len(nums)
	if lenght == 0 {
		return 0
	}
	if lenght == 1 {
		return 1
	}
	slow := 1
	for fast := 1; fast < lenght; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
