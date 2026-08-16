func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		if binarySearch(row, target) {
			return true
		}
	}

	return false
}


func binarySearch(nums []int, target int) bool {
	l, r := 0, len(nums) - 1

	for l <= r {
		m := (l+r) / 2

		if nums[m] == target {
			return true
		}

		if nums[m] > target {
			r = m - 1
		} else {
			l = m+ 1
		}
	}

	return false
}