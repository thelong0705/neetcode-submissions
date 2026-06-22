func groupAnagrams(strs []string) [][]string {
    m := map[[26]int][]string{}
    res := [][]string{}

    for _, str := range strs {
        a := convertToArray(str)
        m[a] = append(m[a], str)
    }

    for _, v := range m {
        res = append(res, v)
    }

    return res
}

func convertToArray(str string) [26]int{
    a := [26]int{}
    for i, _ := range str {
        c := str[i]
        a[c-'a']++
    }
    return a
}