package leetcode

func mergeArray(nums1 []int, m int, nums2 []int, n int) {
	maps := make([]int, 0, m+n)
	p1, p2 := 0, 0

	for {
		if p1 == m {
			maps = append(maps, nums2[p2:]...)
			break
		}
		if p2 == n {
			maps = append(maps, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			maps = append(maps, nums1[p1])
			p1++
		} else {
			maps = append(maps, nums2[p2])
			p2++
		}
	}
	copy(nums1, maps)
}
