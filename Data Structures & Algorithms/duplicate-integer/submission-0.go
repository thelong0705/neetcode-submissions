func hasDuplicate(nums []int) bool {
    exist := map[int]bool{}
    for _, n := range nums {
        if exist[n] {
           return true
        }
        exist[n] = true
    }

    return false
}
