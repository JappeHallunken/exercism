package knapsack

type Item struct {
	Weight, Value int
}

func Knapsack(maxWeight int, items []Item) int {
	dp := make([]int, maxWeight+1)

	for _, item := range items {
		for w := maxWeight; w >= item.Weight; w-- {
			dp[w] = max(dp[w], dp[w-item.Weight]+item.Value)
		}
	}
	return dp[maxWeight]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
