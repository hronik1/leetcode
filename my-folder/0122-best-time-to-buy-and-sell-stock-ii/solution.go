func maxProfit(prices []int) int {
    profit := 0
    nextPrice := prices[len(prices) - 1]
    for i := len(prices) - 2; i >= 0; i -= 1 {
        if dailyProfit := nextPrice - prices[i]; dailyProfit > 0 {
            profit += dailyProfit
        }

        nextPrice = prices[i]
    }

    return profit
}
