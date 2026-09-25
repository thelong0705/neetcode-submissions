// abcabcbb

// abba

func lengthOfLongestSubstring(s string) int {
	var m,l,i int
	

	seen := map[byte]int{}
	
	for i = 0; i < len(s); i++  {
		if idx, ok := seen[s[i]]; ok && idx >= l {
			l = idx+1
		} 
		
		seen[s[i]] = i
		m = max(m, i-l+1)
		
	}



	return m
}
