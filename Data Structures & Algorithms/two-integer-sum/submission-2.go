func twoSum(nums []int, target int) []int {
    m := map[int]int{}

    for i, n := range nums {
        x := target - n
        if idx, ok := m[x]; ok {
            return []int{idx, i}
        }
        m[n] = i
    }

    return []int{}
}