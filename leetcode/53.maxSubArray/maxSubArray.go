package leetcode

func maxSubArray(nums []int) int {
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		//若前面一个元素大于0，则将其加到当前元素上
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
