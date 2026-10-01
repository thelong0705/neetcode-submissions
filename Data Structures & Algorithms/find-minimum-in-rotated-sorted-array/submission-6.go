func findMin(nums []int) int {
	l, r := 0, len(nums) - 1
	mid := (l+r)/2
	m := nums[mid]

	for (l <= r) {
		mid = (l+r)/2
		m  = min(m, nums[mid])
	
		if nums[mid] > nums[r] {
			l = mid+1
		} else {
			r = mid-1
		}
	}

	return m
	
}
