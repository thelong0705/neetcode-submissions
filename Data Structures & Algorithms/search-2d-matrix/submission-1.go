func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix[0])
	l, r := 0, len(matrix) - 1
	
	for l <= r {
		mid := (l+r) / 2

		if matrix[mid][0] <= target && target <= matrix[mid][n-1] {
			return binarySearch(matrix[mid], target)	
		}

		if matrix[mid][0] > target {
			r = mid - 1
		} else if matrix[mid][n-1] < target {
			l = mid + 1
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