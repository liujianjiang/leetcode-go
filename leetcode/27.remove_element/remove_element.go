package leetcode

//给定一个数组 nums 和一个数值 val，将数组中所有等于 val 的元素删除，并返回剩余的元素个数。

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}
