import "slices"

func threeSum(nums []int) [][]int {
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		target := 0 - nums[i]
		r := twoSum(target, nums[i+1:])
		if len(r) > 0 {
			for _, pair := range r {
				pair = append(pair, nums[i])
				slices.Sort(pair)

				seen := false
				for _, x := range res {
					if twoSliceEqual(x, pair) {
						seen = true
					}
				}
				if !seen {
					res = append(res, pair)
				}
			}
		} 
	}

	return res
}

func twoSum(target int, nums[]int) [][]int {
	res := [][]int{}
	seen := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if seen[target-nums[i]] {
			res = append(res, []int{target-nums[i], nums[i]})
		}
		seen[nums[i]] = true
	} 
	return res
}

func twoSliceEqual(s1, s2 []int) bool {
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

