package leetcode

func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := (left + right) / 2
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

//二分查找
func search(nums []int, target int, left int, rifht int) int {
	if left > rifht {
		return -1
	}
	mid := (left + rifht) / 2
	if target > nums[mid] {
		return search(nums, target, mid+1, rifht)
	} else if target < nums[mid] {
		return search(nums, target, left, mid-1)
	} else if target == nums[mid] {
		return mid
	}
	return mid
}
