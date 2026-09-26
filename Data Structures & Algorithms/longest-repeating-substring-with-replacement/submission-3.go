// AABAABBBBBBB
// k = 2

func characterReplacement(s string, k int) int {
	freq, l, maxFreq, m := map[byte]int{}, 0, 0, 0

	for r := 0; r < len(s); r++ {
		freq[s[r]]++
		maxFreq = max(freq[s[r]], maxFreq)
		if (r-l+1) - maxFreq > k {
			freq[s[l]]--
			l++
		} 
		m = max(m, r-l+1)
	}

	return m
}
