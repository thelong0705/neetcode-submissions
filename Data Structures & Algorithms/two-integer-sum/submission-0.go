func twoSum(nums []int, target int) []int {
    m := map[int]int{}

    for i, n := range nums {
        if idx, ok := m[n]; ok && n == target/2 {
            return []int{idx, i}
        }
        m[n] = i

    }

    for i, n := range nums {
        x := target - n
        if idx, ok := m[x]; ok && idx != i {
            return []int{i, idx}
        }
    }

    return []int{}
}