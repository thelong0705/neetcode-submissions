func productExceptSelf(nums []int) []int {
	l, r, n := map[int]int{}, map[int]int{}, len(nums)
	res := make([]int, n)
	l[0], r[n-1]= 1,1

	for i := 1; i < n-1; i++ {
		j := n - 1- i

		
		l[i] = nums[i-1] * l[i-1] 
		r[j] = nums[j+1] * r[j+1]
	}

	for i := 1; i < n-1; i++ {
		res[i] = l[i] * r[i]
	}

	res[0] = nums[1] * r[1]
	res[n-1] = nums[n-2] * l[n-2]

	return res
}
