func maxProfit(prices []int) int {
	l, r, m := 0, 0, 0

	for r < len(prices) {
		if prices[r] < prices[l] {
			l = r
			r = l+1
		} else {
			profit := prices[r] - prices[l]
			if profit > m {
				m = profit
			} else {
				r++
			}
		}
	}

	return m
}
