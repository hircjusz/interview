package besttimetobuysellstock

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxProfit(prices []int) int {

	var ans int
	var lowest int
	lowest = prices[0]

	for _, price := range prices {
		if price < lowest {
			lowest = price
		}
		ans = max(ans, price-lowest)

	}
	return ans
}

//7,1,5,3,6,4
