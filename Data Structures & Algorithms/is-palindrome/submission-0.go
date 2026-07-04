func isPalindrome(s string) bool {
	left, right := 0, len(s) - 1
	for left = 0; left < right; {

		if !isLetterAndNumber(s[left]) {
			    left++
				continue
		}
		for !isLetterAndNumber(s[right]) {
			right--
		}

		if (left > right) || (toLower(s[left]) != toLower(s[right])) {
			return false
		}
		left++
		right--
	}

	return true
}

func isLetterAndNumber(c byte) bool {
	return ('a' <= c && c <= 'z' ) || ('A' <= c && c <= 'Z' ) || ('0' <= c && c <= '9' )
}

func toLower(c byte) byte {
	if ('A' <= c && c <= 'Z' ) {
		return c + 'a' - 'A'
	}
	return c
}