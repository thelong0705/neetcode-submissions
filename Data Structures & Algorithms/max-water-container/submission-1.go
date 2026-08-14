func maxArea(heights []int) int {
	h := heights
	l, r := 0, len(h) - 1
	m := 0
	for l < r {
		m = max(m, min(h[l], h[r]) * (r-l))
		if h[l] < h[r] {
			l++
		} else {
			r--
		}
	}

	return m
}

func max(x,y int) int {
	if x >= y {
		return x
	}
	return y
}

func min(x,y int) int {
	if x <= y {
		return x
	}
	return y
}
