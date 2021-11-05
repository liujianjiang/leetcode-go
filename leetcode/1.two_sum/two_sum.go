package leetcode

func twoSum(nums []int, target int) []int {
	maps := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		other := target - nums[i]

		if _, ok := maps[other]; ok {
			return []int{maps[other], i}
		}

		maps[nums[i]] = i
	}
	return nil
}
